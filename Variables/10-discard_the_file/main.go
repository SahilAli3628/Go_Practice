// ---------------------------------------------------------
// EXERCISE: Discard The File
//
//  1. Print only the directory using `path.Split`
//
//  2. Discard the file part
//
// RESTRICTION
//  Use short declaration
//
// EXPECTED OUTPUT
//  secret/
// ---------------------------------------------------------

package main

import (
	"fmt"
	"path"
)

func main() {

	var directory string
	directory, _ = path.Split("secret/file.txt")
	fmt.Println(directory)

}
