package files_test

import (
	"errors"
	"testing"
	"totality/users/internal/files"
)

func TestReadFromFile(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name string
		in   string
		err  error
	}{
		{
			name: "Ok",
			in:   "testdata/valid.json",
		},
		{
			name: "empty/Ok",
			in:   "testdata/empty.json",
		},
		{
			name: "incorrect/FilePath",
			in:   "incorrect/file/path.json",
			err:  errors.New("no such file or directory"),
		},
	}

	for _, tc := range testcases {
		var tc = tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var got, err = files.ReadFromFile(tc.in)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got %v want %v", err, tc.err)
			}

			if tc.err != nil {
				// If error is expected,then skip further test cases validation.
				return
			}

			if len(got) <= 0 {
				t.Fatalf("failed to read data from file")
			}
		})
	}
}
