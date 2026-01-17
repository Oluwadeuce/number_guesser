package game

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // This means (seed the random generator) Idk what it means yet, but I'm guessing it is a goroutine or something similar
}

func GuessGame(guess int) string {
	correctAns := rand.Intn(10) + 1 //This means to generate a random number

	// fmt.Println("Number Guesser!")
	// for {
	// 	fmt.Println("What is the number")

	// 	var guess int

	// 	_, err := fmt.Scan(&guess)

	 // if err != nil {
	// 	fmt.Println(err)
	// 	break
	// }

	if guess > correctAns {
		return fmt.Sprintf("%v is greater than the number", guess)
	} else if guess < correctAns {
		return fmt.Sprintf("%v is lesser than the number", guess)
	} else {
		return "Correct!"
	}
}
