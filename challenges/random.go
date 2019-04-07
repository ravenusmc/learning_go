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
  var guess int
  fmt.Println("Please guess a number between 1 and 30:")
  fmt.Scan(&guess)
  for number != guess {
    fmt.Println(number)
    fmt.Println("The Guess:", guess)
    if number == guess  {
      fmt.Println("You Guessed the Number!")
    }else if number > guess {
      fmt.Println("You guessed to low!")
    }else if number < guess {
      fmt.Println("You guessed to high!")
    }
    fmt.Println("Please guess a number between 1 and 30:")
    fmt.Scan(&guess)
  }
}

//This function will check the guess of the player
func checkGuess(number int, guess int) {

}
