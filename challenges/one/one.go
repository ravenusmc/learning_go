//Create a program that prompts for your name and than spits out a greeting
//with your name

package main
import (
  "fmt"
)

//First program - simply ask for a name.

// func main() {
//   var name string
//   fmt.Println("What is your name?")
//   fmt.Scan(&name)
//
//   fmt.Println("Good day", name)
// }

//This second program will ask for a name and last name and then ask for a greeting
func main() {
  var name string
  var lastName string
  fmt.Println("What is your name?")
  fmt.Scan(&name)
  fmt.Println("What is your last name?")
  fmt.Scan(&lastName)
  fmt.Println("Hello", name, lastName)
}
