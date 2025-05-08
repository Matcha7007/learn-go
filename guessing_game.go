package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// welcome message
	fmt.Println("Welcome to the Guessing Number Game!")
	fmt.Println("I have choosen a number between 1 and 100.")
	fmt.Println("Can you guess wht it is?")

	var guess int // will be stored guessing number from console 
	for {
		fmt.Println("Enter your guess:")
		fmt.Scanln(&guess) // use addressing & to store the new value from console to guess variable

		// check if the guess is correct
		if guess == target {
			fmt.Println("Congratulation! You guessed the correct number!")
			break
		} else if guess < target {
			fmt.Println("Too low! Try guessing a higher number.")
		} else {
			fmt.Println("Too high! Try guessing a lower number.")
		}
	}
}