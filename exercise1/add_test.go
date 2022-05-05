package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(4, -10)
	want := -6

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
