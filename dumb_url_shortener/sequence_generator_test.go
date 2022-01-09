package main

import "testing"

func TestGenerateSequence(t *testing.T) {
	t.Run("returns a random string of 6 characters", func(t *testing.T) {
		sg := SequenceGenerator{}
		got := len(sg.GenerateSequence())
		want := 6

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
