package main

import (
	"bytes"
	"fmt"
	"testing"
)

type Case struct {
	Description string
	Pancakes    string
	Flips       int
}

func Cases() []Case {
	return []Case{
		Case{
			Description: "one upside down pancake",
			Pancakes:    "-",
			Flips:       1,
		},
		Case{
			Description: "two pancakes, top upside down",
			Pancakes:    "-+",
			Flips:       1,
		},
		Case{
			Description: "two pancakes",
			Pancakes:    "+-",
			Flips:       2,
		},
		Case{
			Description: "all happy pancakes",
			Pancakes:    "+++",
			Flips:       0,
		},
		Case{
			Description: "one happy in the middle",
			Pancakes:    "--+-",
			Flips:       3,
		},
	}
}

func TestFlipPancakes(t *testing.T) {
	tests := []struct {
		In       []byte
		Expected []byte
		Error    error
	}{
		{[]byte("-+-"), []byte("+-+"), nil},
		{[]byte("---"), []byte("+++"), nil},
		{[]byte("+++"), []byte("---"), nil},
		{[]byte("+++"), []byte("---"), nil},
		//{[]byte("5"), []byte("5"), fmt.Errorf("Invalid pancake representation: %c", '5')},
	}

	for _, test := range tests {
		// deep copy original since this is self-mutating
		original := make([]byte, len(test.In))
		copy(original, test.In)

		err := FlipPancakes(test.In)
		equal := bytes.Equal(test.In, test.Expected)

		if err != test.Error {
			t.Fatalf("expected error '%s', got '%s'", test.Error, err)
		}

		if !equal {
			t.Fatalf("expected %s, got %s", test.Expected, test.In)
		}
	}

}

func TestHappyFlips(t *testing.T) {
	for _, test := range Cases() {
		actual := HappyFlips(test.Pancakes)
		if actual != test.Flips {
			t.Fatalf("expected to flip %s %d times, got %d", test.Pancakes, test.Flips, actual)
		}
		fmt.Println("---")
	}
}
