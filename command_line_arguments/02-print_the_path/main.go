// ---------------------------------------------------------
// EXERCISE: Print the Path
//
//  Print the path of the running program by getting it
//  from `os.Args` variable.
//
// HINT
//  Use `go build` to build your program.
//  Then run it using the compiled executable program file.
//
// EXPECTED OUTPUT SHOULD INCLUDE THIS
//  myprogram
//
// IMP: go build -o myprogram main.go
// ---------------------------------------------------------

package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Args[0])

}
