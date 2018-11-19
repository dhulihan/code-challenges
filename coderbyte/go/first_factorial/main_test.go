package main

import "testing"

func TestFirstFactorial(t *testing.T) {
	in := 8
	out := FirstFactorial(in)
	expected := 40320

	if out != expected {
		t.Errorf("wanted %d, got %d", expected, out)
	}
}
