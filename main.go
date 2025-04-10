package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	min := 1
	max := 10

	target := rand.Intn(max-min+1) + min
	attempts := 3

	fmt.Println("Welcome to the number guessing game!")
	// Prompt the user for input
	var input int

	fmt.Print("Enter a number between 1 and 10: ")

	inputhandling(input, target, attempts)
}

func inputhandling(input int, target int, attempts int) {
	// Loop until we get valid input
	for {
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			fmt.Print("Enter a number between 1 and 10: ")
			continue
		}

		if input < 1 || input > 10 {
			fmt.Println("Please select a number between 1 and 10")
			fmt.Print("Enter a number between 1 and 10: ")
			continue
		}

		// If we reach here, input is valid, break the loop
		break
	}

	// Proceed with the game logic (using valid input)
	if input == target {
		fmt.Println("Congratulations! You guessed the number.")
		return
	} else {
		attempts--

		if input > target {
			fmt.Println("Too high!")
		}

		if input < target {
			fmt.Println("Too low!")
		}

		if attempts == 0 {
			fmt.Println("Sorry, you've used all your attempts. The number was:", target)
			return
		} else {
			fmt.Printf("You have %d attempts left.\n", attempts)
			fmt.Print("Enter a number between 1 and 10: ")
			inputhandling(input, target, attempts)
		}
	}
}
