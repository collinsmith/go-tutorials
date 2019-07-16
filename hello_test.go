package main

import "testing"

func TestHello(t *testing.T) {
  assertMessage := func(t *testing.T, got, want string) {
    t.Helper()
    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  }

  t.Run("saying hello to people", func(t *testing.T) {
    got := Hello("Collin")
    want := "Hello, Collin"
    assertMessage(t, got, want)
  })

  t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
    got := Hello("")
    want := "Hello, world"
    assertMessage(t, got, want)
  })
}