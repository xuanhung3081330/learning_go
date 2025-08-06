package main

import "testing"

func TestHello(t *testing.T){
	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T){
		got := Hello("Hung", "French")
		want := "Bonjour, Hung"
		assertCorrectMessage(t, got, want)
	})
}

func TestRepeat(t *testing.T){
	got := Repeat("a", 5)
	want := "aaaaa"
	assertCorrectMessage(t, got, want)
}

func BenchmarkRepeat(b *testing.B){
	for b.Loop(){
		Repeat("a", 100)
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
 	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
