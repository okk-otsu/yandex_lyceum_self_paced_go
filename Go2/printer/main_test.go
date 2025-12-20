package printer

import (
	"errors"
	"testing"
	"unicode/utf8"
)

var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}

	return utf8.RuneCount(input), nil
}

func TestUTF8(t *testing.T) {
	cases := []struct {
		name  string
		value []byte
		want  int
		err   error
	}{
		// тестовые данные № 1
		{
			name:  "number 1",
			value: []byte{'a', 'b', 'c', 'd'},
			want:  4,
			err:   nil,
		},
		// тестовые данные № 2
		{
			name:  "number 2",
			value: []byte{'a', 'b', 'c'},
			want:  3,
			err:   nil,
		},
		// тестовые данные № 3
		{
			name:  "number 3",
			value: []byte{0xff, 0xfe},
			want:  0,
			err:   ErrInvalidUTF8,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			got, err := GetUTFLength(tc.value)
			if err != tc.err {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}
