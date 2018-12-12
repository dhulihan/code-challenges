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

	// DebugMode determines if we should print debug information
	DebugMode = true
)

// FlipPancake simply inverts a pancake
func FlipPancake(pancake byte) (byte, error) {
	switch pancake {
	case Happy:
		return Sad, nil
	case Sad:
		return Happy, nil
	default:
		debugf("Got an error during FlipPancake: %s\n", "INVALID CHAR")
		return '?', fmt.Errorf("Invalid pancake representation: %c", pancake)
	}
}

// FlipPancakes performs an in-place modification of slice
func FlipPancakes(pancakes []byte) error {
	for i := range pancakes {
		ret, err := FlipPancake(pancakes[i])

		if err != nil {
			debugf("Got an error during FlipPancakes: %s\n", err)
			return err
		}

		pancakes[i] = ret
	}

	return nil
}

// HappyFlips returns the number of flips needed to flip pancakes
// so happy face (+) is right side up. Reads left to right.
func HappyFlips(pancakes string) (int, error) {
	chars := []byte(pancakes)

	// position of the flip
	var flips int
	var nextPancake byte

TraversePancakes:
	for {
		debugf("\nGiven %s\n", chars)
		for i := range chars {

			if i < len(chars)-1 {
				nextPancake = chars[i+1]
			} else {
				nextPancake = Happy
			}

			debugf("i: %d nextPancake: %c, currPancake: %c\n", i, nextPancake, chars[i])
			// if we hit a sad pancake
			if chars[i] == Sad && nextPancake == Happy {
				debugf("Stopping at i: %d, flips: %d\n", i, flips)

				// Determine end of subslice to flip
				var end int
				if len(chars) <= 1 {
					end = len(chars)
				} else {
					end = i + 1
				}

				debugf("Flipping %s\n", chars[:end])
				err := FlipPancakes(chars[:end])

				if err != nil {
					debugf("Got an error: %s\n", err)
					return -1, err
				}

				// reset
				flips++
				debugf("Trying again with %s\n", chars)
				continue TraversePancakes
			}
		}
		break
	}

	return flips, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n := 1

	for scanner.Scan() {
		text := scanner.Text()

		flips, err := HappyFlips(text)

		if err != nil {
			debugf("Got an error: %s\n", err)
			continue
		}

		fmt.Printf("Case #%d: %d\n", n, flips)
		n++
	}
}

// cute debug function
func debugf(format string, vals ...interface{}) {
	if DebugMode {
		fmt.Printf(format, vals...)
	}
}
