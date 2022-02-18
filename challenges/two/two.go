// This assigment asked to make a program that will take in a person's name and then
// calculate its length

package main

import (
	"fmt"
)

func main() {

	var name string
	fmt.Println("What is your name?")
	fmt.Scan(&name)
	length := len(name)
	fmt.Println(length)

}
