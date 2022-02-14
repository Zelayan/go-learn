package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(got string, want string, t *testing.T) {
		// tell test is helper test
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	

	t.Run("say hello to people",func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello,Chris"
	
		assertCorrectMessage(got, want, t)
	})


	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello,world"
		assertCorrectMessage(got, want, t)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hello,Elodie"
		assertCorrectMessage(go)
	})
}
