package main

import (
	"fmt"
	"os"
)

// Errors

/*
* Go programs are not built around with exceptions. its
* a bad idea to use exeception.
 */

// Built in type error handling
func printer(msg string) error {
	_, err := fmt.Printf("%s \n", msg)

	return err
}

func main() {
	if err := printer("Hello, world!"); err != nil {
		os.Exit(1)
	}
}
