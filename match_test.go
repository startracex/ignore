package ignore

import (
	"testing"
)

func TestMatch(t *testing.T) {
	list := []struct {
		glob  string
		path  string
		match bool
	}{
		{"a", "a", true},
		{"a", "b/a", true},
		{"a", "b/a", true},
		{"/a", "a", true},
		{"/a", "b/a", false},
		{"!a", "a", false},
		{"!a", "b", true},

		{"!/a", "a", false},
		{"/!a", "a", false},
		{"**/*.a", "b/a.a", true},
		{"/**", "a/", true},
		{"/**/*.a", "b/a", false},
	}
	for _, l := range list {
		if match := Match(l.glob, l.path); match != l.match {
			t.Errorf("Match(%q, %q) = %v, want %v", l.glob, l.path, match, l.match)
		}
	}
}
