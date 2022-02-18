package main

import "fmt"

func main() {

	var firstName string
	var lastName string
	
  fmt.Println("What is your first name?")
	fmt.Scan(&firstName)
	
	fmt.Println("What is your last name?")
	fmt.Scan(&lastName)

	fmt.Println("Hello ", firstName, lastName)
}