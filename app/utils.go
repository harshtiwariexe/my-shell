package main

import (
	"os"
	"path/filepath"
	"strings"
)

func findExecutable(cmd string) string {
	pathEnv := os.Getenv("PATH")
	parts := strings.Split(pathEnv, ":")

	for _, dir := range parts {
		if dir == "" {
			dir = "."
		}
		candidate := filepath.Join(dir, cmd)

		fi, err := os.Stat(candidate)

		if err != nil {
			continue
		}
		if fi.IsDir() {
			continue
		}
		if fi.Mode().Perm()&0111 != 0 {
			return candidate
		}
	}
	return ""

}
