package main

import (
	"encoding/json"
	"net/http"
	"numberguesser/game"
	"strconv"
)

type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	guessStr := r.URL.Query().Get("number")  // This gets the "number" parameter from the URL
	guess, err := strconv.Atoi(guessStr)  //This converts it to an integer

	if err != nil {
		json.NewEncoder(w).Encode(Response{"Please enter a valid number."})
		return
	}

	message := game.GuessGame(guess)   //This calls the game logic from game.go
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{message})
}

func main() {
	http.HandleFunc("/guess", handler)
	http.ListenAndServe(":8080", nil)
}
