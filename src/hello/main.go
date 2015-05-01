package main

import (
	"errors"
	"fmt"
	"os"
)

// Import errors pacakge and declare a string variable with the error msg
var (
	errorEmptyString = errors.New("Unwilling to print an empty string")
)

// Errors

/*
* Go programs are not built around with exceptions. its
* a bad idea to use exeception.
*
* For custom errors, we can use fmt package "Errorf" function
* Another way to defined errors is to use Error package. the advantage here is
* we can assigned what type of error etc..
*
 */

// Built in type error handling
func printer(msg string) error {
	// Using Fmt.Errorf to give custom error msg
	if msg == "" {
		// use the error variable declared
		return errorEmptyString
	}
	_, err := fmt.Printf("%s \n", msg)

	return err
}

func main() {
	if err := printer(""); err != nil {
		// since we used error pcakge, now we can compare the error msg
		if err == errorEmptyString {
			fmt.Printf("You rited to print a empty string!")
		} else {
			fmt.Printf("printer failed: %s\n", err)
		}
		os.Exit(1)
	}
}
