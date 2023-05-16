package osutil_test

import (
	"errors"
	"os"
	"testing"

	"git.sr.ht/~jamesponddotco/janus/internal/osutil"
)

func TestReadFile(t *testing.T) {
	t.Parallel()

	tempFile, err := os.CreateTemp(os.TempDir(), "prefix")
	if err != nil {
		t.Fatalf("Error creating temporary file, %s", err)
	}

	defer t.Cleanup(func() {
		if err = tempFile.Close(); err != nil {
			t.Fatalf("Error closing temporary file, %s", err)
		}
	})

	text := "Hello, World!"

	if _, err = tempFile.WriteString(text); err != nil {
		t.Fatalf("Error writing to temporary file, %s", err)
	}

	tests := []struct {
		name     string
		filename string
		want     string
		wantErr  bool
	}{
		{
			name:     "File exists",
			filename: tempFile.Name(),
			want:     text,
			wantErr:  false,
		},
		{
			name:     "File does not exist",
			filename: "non-existent-file.txt",
			want:     "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := osutil.ReadFile(tt.filename)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected an error but got none")
				} else if !errors.Is(err, os.ErrNotExist) {
					t.Errorf("expected 'file does not exist' error but got %s", err)
				}
			} else if err != nil {
				t.Errorf("got unexpected error: %v", err)
			}

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
