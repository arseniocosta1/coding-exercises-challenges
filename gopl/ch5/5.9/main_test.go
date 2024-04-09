package main

import (
	"strings"
	"testing"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TestExpand(t *testing.T) {
	titleCaser := cases.Title(language.English)
	for _, test := range []struct {
		s, want string
		f       func(string) string
	}{
		{"$foo", "FOO", strings.ToTitle},
		{"$FOO$BaR", "foobar", strings.ToLower},
		{"hello $foo, this is your $bAr", "hello Foo, this is your Bar",
			func(s string) string {
				return titleCaser.String(s)
			},
		},
	} {
		got := expand(test.s, test.f)
		if got != test.want {
			t.Errorf("expand(%q, f) = %q, want %q", test.s, got, test.want)
		}
	}
}

func BenchmarkExpand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		expand("hello $foo, this is your $bAr", strings.ToUpper)
	}
}
