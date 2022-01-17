package mutils

import "github.com/mocaberos/glibs/mfs"

// IsRunningOnDocker returns true if this program is running on docker.
func IsRunningOnDocker() bool {
	return mfs.IsFileExists("/.dockerenv")
}
