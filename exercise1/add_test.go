package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(5, 4, 2, -10, 4.2, "l")
	want := 1

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
