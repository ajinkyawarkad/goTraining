package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 4)
	want := 6

	if got != want {
		t.Errorf("failed the test!")

	}
}
