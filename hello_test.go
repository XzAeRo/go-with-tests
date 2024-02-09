package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Victor")
	want := "Hello, Victor"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// The convention is to name the test file the same as the source file but with _test.go added to the end.
