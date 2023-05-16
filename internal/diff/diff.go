// Package diff provides a diff implementation for the diff command.
package diff

import (
	"fmt"
	"strings"

	"git.sr.ht/~jamesponddotco/janus/internal/osutil"
	"github.com/aymanbagabas/go-udiff"
	"github.com/aymanbagabas/go-udiff/myers"
	"github.com/fatih/color"
)

// Options represents the options for the diff command.
type Options struct {
	// Colorize is whether to colorize the output.
	Colorize bool
}

// Compare compares two files and returns the differences between them given the
// provided options.
func Compare(oldFile, newFile string, opts *Options) (string, error) {
	oldContent, err := osutil.ReadFile(oldFile)
	if err != nil {
		return "", fmt.Errorf("failed to read old file: %w", err)
	}

	newContent, err := osutil.ReadFile(newFile)
	if err != nil {
		return "", fmt.Errorf("failed to read new file: %w", err)
	}

	edits := myers.ComputeEdits(oldContent, newContent)

	diff, err := udiff.ToUnified(oldFile, newFile, oldContent, edits)
	if err != nil {
		return "", fmt.Errorf("failed to generate unified diff: %w", err)
	}

	if opts.Colorize {
		var builder strings.Builder

		builder.Grow(len(diff))

		for _, line := range strings.Split(diff, "\n") {
			switch {
			case strings.HasPrefix(line, "-"):
				builder.WriteString(color.RedString(line))
			case strings.HasPrefix(line, "+"):
				builder.WriteString(color.GreenString(line))
			default:
				builder.WriteString(line)
			}

			builder.WriteString("\n")
		}

		colorizedDiff := strings.TrimSuffix(builder.String(), "\n")

		return colorizedDiff, nil
	}

	return diff, nil
}
