package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Generate a random number between 1 and 6 (inclusive)
	randomNumber := rand.Intn(6) + 1

	// Print the random number
	fmt.Println("Random number between 1 and 6:", randomNumber)

	// Switch case to handle different values
	switch randomNumber {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
		// fmt.Println("The number 3 is considered lucky in many cultures.")
		// fallthrough
	case 4:
		fmt.Println("Four")
		fmt.Println("The number 4 is considered unlucky in some East Asian cultures.")
		fallthrough
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six")
	default:
		fmt.Println("Unknown number")
	}
}
