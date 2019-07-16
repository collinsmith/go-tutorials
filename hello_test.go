package main

import "testing"

func TestHello(t *testing.T) {
  got := Hello("Collin")
  want := "Hello, Collin"

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}