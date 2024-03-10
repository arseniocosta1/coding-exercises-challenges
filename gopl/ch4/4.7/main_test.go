package main

import (
	"bytes"
	"testing"
)

func TestRevUTF8(t *testing.T) {
	s := []byte("\u00A0Räksmörgås")

	reverse(s)
	want := []byte("sågrömskäR\u00A0")
	if bytes.Compare(s, want) != 0 {
		t.Errorf("got %v, want %v", string(s), want)
	}
}

func BenchmarkRevUTF8(b *testing.B) {
	s := []byte("Large string with a lot of characters. 🤔👋🏿👋🏿👋🏿👋🏿👋🏿👋🏿👋🏿👋🏿👋🏿👋🏿👋🏿👋🏿👋🏿 showing that no allocation occurs")
	for i := 0; i < b.N; i++ {
		reverse(s)
	}
}
