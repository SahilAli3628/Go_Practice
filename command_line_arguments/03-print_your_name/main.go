// ---------------------------------------------------------
// EXERCISE: Print Your Name
//
//  Get it from the first command-line argument
//
// INPUT
//  Call the program using your name
//
// EXPECTED OUTPUT
//  It should print your name
//
// EXAMPLE
//  go run main.go inanc
//
//    inanc
//
// BONUS: Make the output like this:
//
//  go run main.go inanc
//    Hi inanc
//    How are you?
// ---------------------------------------------------------

package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Args[1])

	//Bonus
	fmt.Printf("Hi %s\nHow are you?", os.Args[1])

}
