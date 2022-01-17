package mfs

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// IsPathExists reports whether the file or directory exists.
func IsPathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsFileExists reports whether the file exists.
func IsFileExists(path string) bool {
	return IsFile(path)
}

// IsDir reports whether the target path is a directory.
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir()
	} else {
		return false
	}
}

// IsFile reports whether the target path is a file.
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err == nil {
		return !info.IsDir()
	} else {
		return false
	}
}

// SelfDir returns The directory that contains the executable file.
// If current executable is a symbolic link,
// Returns the real path to the directory that contains the executable.
func SelfDir() (path string, err error) {
	if path, err = os.Executable(); err != nil {
		return
	}
	if path, err = filepath.EvalSymlinks(path); err != nil {
		return
	}
	if path, err = filepath.Abs(path); err != nil {
		return
	}
	return filepath.Dir(path), nil
}

// GetParentDir returns the path of parent directory.
func GetParentDir(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(absPath, "/")
	if i == -1 {
		return absPath, nil
	} else if i == 0 {
		return "/", nil
	} else {
		return absPath[:i], nil
	}
}

// CopyFile from `source` to `dest`
func CopyFile(source, dest string) (err error) {
	var fd1, fd2 *os.File
	if fd1, err = os.Open(source); err != nil {
		return err
	}
	defer func(fd1 *os.File) {
		if e := fd1.Close(); e != nil {
			err = e
		}
	}(fd1)
	if fd2, err = os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644); err != nil {
		return err
	}
	defer func(fd1 *os.File) {
		if e := fd1.Close(); e != nil {
			err = e
		}
	}(fd1)
	_, err = io.Copy(fd2, fd1)
	return err
}

// GetFileSize return the size of the file as bytes.
func GetFileSize(filename string) (int64, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return 0, err
	}
	return info.Size(), nil
}
