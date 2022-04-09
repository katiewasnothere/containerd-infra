/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package v2

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/containerd/containerd/identifiers"
	"github.com/containerd/containerd/log"
	"github.com/containerd/containerd/mount"
	"github.com/containerd/containerd/namespaces"
	cioutil "github.com/containerd/containerd/pkg/ioutil"
)

const configFilename = "config.json"

// LoadBundle loads an existing bundle from disk
func LoadBundle(ctx context.Context, root, id string) (*Bundle, error) {
	ns, err := namespaces.NamespaceRequired(ctx)
	if err != nil {
		return nil, err
	}
	return &Bundle{
		ID:        id,
		Path:      filepath.Join(root, ns, id),
		Namespace: ns,
	}, nil
}

// NewBundle returns a new bundle on disk
func NewBundle(ctx context.Context, root, state, id string, spec []byte) (b *Bundle, err error) {
	if err := identifiers.Validate(id); err != nil {
		return nil, fmt.Errorf("invalid task id %s: %w", id, err)
	}

	ns, err := namespaces.NamespaceRequired(ctx)
	if err != nil {
		return nil, err
	}
	work := filepath.Join(root, ns, id)
	b = &Bundle{
		ID:        id,
		Path:      filepath.Join(state, ns, id),
		Namespace: ns,
	}
	var paths []string
	defer func() {
		if err != nil {
			for _, d := range paths {
				os.RemoveAll(d)
			}
		}
	}()
	// create state directory for the bundle
	if err := os.MkdirAll(filepath.Dir(b.Path), 0711); err != nil {
		return nil, err
	}
	if err := os.Mkdir(b.Path, 0700); err != nil {
		return nil, err
	}
	if err := prepareBundleDirectoryPermissions(b.Path, spec); err != nil {
		return nil, err
	}
	paths = append(paths, b.Path)
	// create working directory for the bundle
	if err := os.MkdirAll(filepath.Dir(work), 0711); err != nil {
		return nil, err
	}
	rootfs := filepath.Join(b.Path, "rootfs")
	if err := os.MkdirAll(rootfs, 0711); err != nil {
		return nil, err
	}
	paths = append(paths, rootfs)
	if err := os.Mkdir(work, 0711); err != nil {
		if !os.IsExist(err) {
			return nil, err
		}
		os.RemoveAll(work)
		if err := os.Mkdir(work, 0711); err != nil {
			return nil, err
		}
	}
	paths = append(paths, work)
	// symlink workdir
	if err := os.Symlink(work, filepath.Join(b.Path, "work")); err != nil {
		return nil, err
	}
	// write the spec to the bundle
	err = os.WriteFile(filepath.Join(b.Path, configFilename), spec, 0666)
	return b, err
}

// Bundle represents an OCI bundle
type Bundle struct {
	// ID of the bundle
	ID string
	// Path to the bundle
	Path string
	// Namespace of the bundle
	Namespace string
}

// Delete a bundle atomically
func (b *Bundle) Delete(ctx context.Context) error {
	work, werr := os.Readlink(filepath.Join(b.Path, "work"))
	rootfs := filepath.Join(b.Path, "rootfs")

	// on windows hcsshim writes panic logs in the bundle directory in a file named
	// "panic.log" log those messages (if any).
	// Read only upto 1MB worth of data from this file. If the file is larger
	// than that, log that.
	readLimit := int64(1024 * 1024) // 1MB
	logBytes, err := cioutil.LimitedRead(filepath.Join(b.Path, "panic.log"), readLimit)
	if err == nil && len(logBytes) > 0 {
		if int64(len(logBytes)) == readLimit {
			log.G(ctx).Warnf("shim panic log file %s is larger than 1MB, logging only first 1MB", filepath.Join(b.Path, "panic.log"))
		}
		log.G(ctx).WithField("log", string(logBytes)).Warn("found shim panic logs during delete")
	} else if err != nil && !os.IsNotExist(err) {
		log.G(ctx).WithError(err).Warn("failed to open shim panic log")
	}

	if err := mount.UnmountAll(rootfs, 0); err != nil {
		return fmt.Errorf("unmount rootfs %s: %w", rootfs, err)
	}
	if err := os.Remove(rootfs); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove bundle rootfs: %w", err)
	}
	err = atomicDelete(b.Path)
	if err == nil {
		if werr == nil {
			return atomicDelete(work)
		}
		return nil
	}
	// error removing the bundle path; still attempt removing work dir
	var err2 error
	if werr == nil {
		err2 = atomicDelete(work)
		if err2 == nil {
			return err
		}
	}
	return fmt.Errorf("failed to remove both bundle and workdir locations: %v: %w", err2, err)
}
