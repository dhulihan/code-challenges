package main

import "fmt"

const (
	HAPPY = '+'
	SAD   = '-'
)

// FlipPancake simply inverts a pancake
func FlipPancake(pancake byte) (byte, error) {
	switch pancake {
	case HAPPY:
		return SAD, nil
	case SAD:
		return HAPPY, nil
	default:
		return '?', fmt.Errorf("Invalid pancake representation: %c", pancake)
	}
}

// FlipPancakes performs an in-place modification of slice
func FlipPancakes(pancakes []byte) error {
	for i := range pancakes {
		ret, err := FlipPancake(pancakes[i])

		if err != nil {
			return err
		}

		pancakes[i] = ret
	}

	//fmt.Printf("Original pancakes: %s, new pancakes: %s\n", pancakes, pancakes)

	return nil
}

// HappyFlips returns the number of flips needed to flip pancakes
// so happy face (+) is right side up. Reads left to right.
func HappyFlips(pancakes string) int {
	chars := []byte(pancakes)

	// position of the flip
	var flips int
	var nextPancake byte

TraversePancakes:
	for {
		fmt.Printf("\nGiven %s\n", chars)
		for i := range chars {

			if i < len(chars)-1 {
				nextPancake = chars[i+1]
			} else {
				nextPancake = HAPPY
			}

			//fmt.Printf("i: %d nextPancake: %c, currPancake: %c\n", i, nextPancake, chars[i])
			// if we hit a sad pancake
			if chars[i] == SAD && nextPancake == HAPPY {
				fmt.Printf("Stopping at i: %d, flips: %d\n", i, flips)

				if len(chars) <= 1 {
					FlipPancakes(chars[:len(chars)])
				} else {
					fmt.Printf("Flipping %s\n", chars[:i+1])
					FlipPancakes(chars[:i+1])
				}

				// reset
				flips++
				fmt.Printf("Trying again with %s\n", chars)
				continue TraversePancakes
			}
		}
		break
	}

	return flips
}

func main() {
	//cases := Cases()
}
