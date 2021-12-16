package files_test

import (
	"testing"
	"totality/users/internal/files"

	"github.com/google/go-cmp/cmp"
)

func TestReadFromFile(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name string
		in   string
		want []byte
		err  error
	}{
		{
			name: "Ok",
			in:   "testdata/valid.json",
			want: []byte(`{
				"id": 1,
				"name": "Steve",
				"city": "LA",
				"phone": 9999999,
				"height": 5.10,
				"married": true
			}`),
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

			if !cmp.Equal(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}
