//The purpose of this program will be to make a simple guessing game.
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  number := getRandomNumber()
  makeGuess(number)
}

//This function will get a random number
func getRandomNumber() int {
  rand.Seed(time.Now().UnixNano())
  number := rand.Intn(30)
  return number
}

//This function will allow the user to make a guess
func makeGuess(number int) {
  var guess string
  fmt.Println("Please guess a number between 1 and 30:")
  fmt.Scan(&guess)
  fmt.Println("You guess is:", guess)
  fmt.Println(number)
}
