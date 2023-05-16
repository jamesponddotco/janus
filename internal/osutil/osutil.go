// Package osutil provides useful functions for working with the operating system.
package osutil

import (
	"fmt"
	"os"

	"git.sr.ht/~jamesponddotco/xstd-go/xunsafe"
)

// ReadFile reads the file named by filename and returns the contents as a string.
func ReadFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return xunsafe.BytesToString(bytes), nil
}
