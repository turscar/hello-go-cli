package main

import "testing"

func TestWorld(t *testing.T) {
	want := "World"
	got := world()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
