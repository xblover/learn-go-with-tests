package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("xubo")
	want := "hello, xubo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
