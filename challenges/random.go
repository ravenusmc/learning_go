//The purpose of this program will be to make a simple guessing game.
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  number := getRandomNumber()
  // fmt.Println(number)
}

func getRandomNumber() int {
  rand.Seed(time.Now().UnixNano())
  number := rand.Intn(30)
  return number
}
