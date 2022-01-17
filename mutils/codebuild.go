package mutils

import "os"

// IsRunningOnCodeBuild returns true if this program is running on codebuild.
func IsRunningOnCodeBuild() bool {
	return os.Getenv("CODEBUILD_BUILD_ID") != ""
}
