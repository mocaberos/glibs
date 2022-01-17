package menv

import "os"

// GetEnv get environment value by key name with default value.
func GetEnv(name string, def ...string) string {
	val := os.Getenv(name)
	if val == "" && len(def) > 0 {
		val = def[0]
	}
	return val
}

// SetEnv can set environment value.
var SetEnv = os.Setenv
