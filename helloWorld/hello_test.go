package main

import "testing"

func TestHello(t *testing.T) {

  assertCorrectMessage := func(t *testing.T, got, want string) {
    t.Helper()
    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  }

  t.Run("saying hello to people", func(t *testing.T) {
    got := Hello("Sean", "English")
    want := "Hello, Sean"
    assertCorrectMessage(t, got, want)
  })

  t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
    got := Hello("", "")
    want := "Hello, World"
    assertCorrectMessage(t, got, want)
  })

  t.Run("in German", func(t *testing.T) {
    got := Hello("Hans", "German")
    want := "Hallo, Hans"
    assertCorrectMessage(t, got, want)
  })

  t.Run("in french", func(t *testing.T) {
    got := Hello("Luc", "French")
    want := "Bonjour, Luc"
    assertCorrectMessage(t, got, want)
  })

  t.Run("in spanish", func(t *testing.T) {
    got := Hello("Carlos", "Spanish")
    want := "Hola, Carlos"
    assertCorrectMessage(t, got, want)
  })
}
