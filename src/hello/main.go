package main

import (
	"fmt"
	"os"
)

// Errors

/*
* Go programs are not built around with exceptions. its
* a bad idea to use exeception.
*
* For custom errors, we can use fmt package "Errorf" function
 */

// Built in type error handling
func printer(msg string) error {
	// Using Fmt.Errorf to give custom error msg
	if msg == "" {
		return fmt.Errorf("Unwilling to print an empty string")
	}
	_, err := fmt.Printf("%s \n", msg)

	return err
}

func main() {
	if err := printer(""); err != nil {
		fmt.Printf("printer failed: %s\n", err)
		os.Exit(1)
	}
}
