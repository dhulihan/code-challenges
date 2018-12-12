package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	// Happy defines the char for a happy pancake
	Happy = '+'

	// Sad defines the char for a sad/upside-down pancake
	Sad = '-'
)

// Check that pancake representation is valid
func validatePancake(pancake byte) error {
	switch pancake {
	case Happy:
		return nil
	case Sad:
		return nil
	default:
		return fmt.Errorf("Invalid pancake representation: %c", pancake)
	}
}

// FlipPancake simply inverts a pancake
func FlipPancake(pancake byte) byte {
	switch pancake {
	case Happy:
		return Sad
	case Sad:
		return Happy
	default:
		panic(fmt.Sprintf("Cannot flip a %c pancake", pancake))
	}
}

// FlipPancakes performs an in-place modification of slice
func FlipPancakes(pancakes []byte) error {
	for i := range pancakes {
		ret := FlipPancake(pancakes[i])
		pancakes[i] = ret
	}

	return nil
}

// HappyFlips returns the number of flips needed to flip pancakes
// so happy face (+) is right side up. Reads left to right. This approach
// modifies the underlying array, so no additional data/copies of the array
// need to be created.
func HappyFlips(stack string) (int, error) {
	pancakes := []byte(stack)

	// position of the flip
	var flips int
	var nextPancake byte
	var err error

TraversePancakes:
	for {
		for i := range pancakes {
			// validate pancake
			err = validatePancake(pancakes[i])
			if err != nil {
				return -1, err
			}

			// Find next pancake. If at the end of the slice, cap it off
			// with a fake happy pancake.
			if i < len(pancakes)-1 {
				nextPancake = pancakes[i+1]
			} else {
				nextPancake = Happy
			}

			// if we hit a sad pancake
			if pancakes[i] == Sad && nextPancake == Happy {
				// Determine subslice range that we'll be flipping
				var end int
				if len(pancakes) <= 1 {
					end = len(pancakes)
				} else {
					end = i + 1
				}

				FlipPancakes(pancakes[:end])

				flips++
				continue TraversePancakes
			}
		}
		break
	}

	return flips, nil
}

func main() {
	// Read pancakes on a per-line basis from stdin.
	scanner := bufio.NewScanner(os.Stdin)
	n := 1

	for scanner.Scan() {
		text := scanner.Text()

		flips, err := HappyFlips(text)

		if err != nil {
			continue
		}

		fmt.Printf("Case #%d: %d\n", n, flips)
		n++
	}
}
