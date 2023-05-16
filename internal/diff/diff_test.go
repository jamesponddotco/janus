package diff_test

import (
	"testing"

	"git.sr.ht/~jamesponddotco/janus/internal/diff"
	"github.com/fatih/color"
)

func TestCompare(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		oldFile     string
		newFile     string
		opts        *diff.Options
		expectedOut string
		wantErr     bool
	}{
		{
			name:        "no_changes",
			oldFile:     "testdata/a.txt",
			newFile:     "testdata/a.txt",
			opts:        &diff.Options{Colorize: false},
			expectedOut: "",
			wantErr:     false,
		},
		{
			name:        "line_change_without_color",
			oldFile:     "testdata/a.txt",
			newFile:     "testdata/b.txt",
			opts:        &diff.Options{Colorize: false},
			expectedOut: "--- testdata/a.txt\n+++ testdata/b.txt\n@@ -1 +1 @@\n-Hello, world!\n+Hello, Go!\n",
			wantErr:     false,
		},
		{
			name:    "line_change_with_color",
			oldFile: "testdata/a.txt",
			newFile: "testdata/b.txt",
			opts:    &diff.Options{Colorize: true},
			expectedOut: color.RedString("--- testdata/a.txt") + "\n" +
				color.GreenString("+++ testdata/b.txt") + "\n" +
				"@@ -1 +1 @@\n" +
				color.RedString("-Hello, world!") + "\n" +
				color.GreenString("+Hello, Go!") + "\n",
			wantErr: false,
		},
		{
			name:        "error_reading_old_file",
			oldFile:     "testdata/non_existent_file.txt",
			newFile:     "testdata/b.txt",
			opts:        &diff.Options{Colorize: false},
			expectedOut: "",
			wantErr:     true,
		},
		{
			name:        "error_reading_new_file",
			oldFile:     "testdata/a.txt",
			newFile:     "testdata/non_existent_file.txt",
			opts:        &diff.Options{Colorize: false},
			expectedOut: "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			out, err := diff.Compare(tt.oldFile, tt.newFile, tt.opts)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Compare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if out != tt.expectedOut {
				t.Errorf("unexpected output: got %q, want %q", out, tt.expectedOut)
			}
		})
	}
}
