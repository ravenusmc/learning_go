//The purpose of this program will be to make a simple guessing game.
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  var number int
  difficulty := selectDifficulty()
  if difficulty == 1 {
    number := getRandomNumber()
    makeGuessEasy(number)
  }else {
    makeGuessHard(number)
  }

}

//This function will allow the user to select the difficulty of the game
func selectDifficulty() int {
  var difficulty int
  fmt.Println("Please select your game difficulty:")
  fmt.Println("1. Easy")
  fmt.Println("2. Hard")
  fmt.Scan(&difficulty)
  return difficulty
}

//This function will get a random number
func getRandomNumber() int {
  rand.Seed(time.Now().UnixNano())
  number := rand.Intn(30)
  return number
}

//This function will allow the user to make a guess
func makeGuessEasy(number int) {
  var guess int
  fmt.Println("Please guess a number between 1 and 30:")
  fmt.Scan(&guess)
  for number != guess {
    if number > guess {
      fmt.Println("You guessed to low!")
    }else if number < guess {
      fmt.Println("You guessed to high!")
    }
    fmt.Println("Please guess a number between 1 and 30:")
    fmt.Scan(&guess)
    if number == guess  {
      fmt.Println("You Guessed the Number!")
    }
  }
}

//This function is for the hard game
func makeGuessHard(number int){
  var guess int
  fmt.Println("Please guess a number between 1 and 100:")
  fmt.Scan(&guess)
  var count = 0
  for count <= 6 {
    if number > guess {
      fmt.Println("You guessed to low!")
    }else if number < guess {
      fmt.Println("You guessed to high!")
    }
    fmt.Println("Please guess a number between 1 and 100:")
    fmt.Scan(&guess)
    if number == guess  {
      fmt.Println("You Guessed the Number!")
      break;
    }
    count++
  }
}







//
