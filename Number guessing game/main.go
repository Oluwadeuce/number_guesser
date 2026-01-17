// INITIAL CODE
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {


	rand.Seed(time.Now().UnixNano()) // This means (seed the random generator) Idk what it means yet, but I'm guessing it is a goroutine or something similar
	correctAns := rand.Intn(1000) + 1  // This generates a random number between 1 and 1000
	fmt.Println("Number Guesser!")
	for {
		fmt.Println("Guess a number between 1 and 1000:")

		var guess int

		_, err := fmt.Scan(&guess)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if guess > correctAns {
			fmt.Printf("%v is greater than the number\n", guess)
		} else if guess < correctAns {
			fmt.Printf("%v is lesser than the number\n", guess)
		} else if guess == correctAns {
			fmt.Println("Correct!")
			break
		}
	}
}


