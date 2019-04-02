//Create a program that prompts for your name and than spits out a greeting
//with your name

package main
import (
  "fmt"
  "math/rand"
)

func main() {
  var name string
  fmt.Println("What is your name?")
  fmt.Scan(&name)

  fmt.Println("Good day", name)

}
