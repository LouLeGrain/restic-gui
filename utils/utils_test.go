package utils

import (
"testing"
)

func TestSetFile(t *testing.T) {
	tests := []struct {
		path, cont, want string
	}{
		{"", "", ""},
		{
			"/Users/andi/test.txt",
			"nice content",
			"/Users/andi/test.txt",
		},
	}
	for _, test := range tests {
		if got := SetFile(test.path, test.cont); test.want != got {
			t.Errorf("%q: want\n\t%q\nbut got\n\t%q", test.path, test.want, got)
		}
	}
}