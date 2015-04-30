package main

import (
	"fmt"
	"os"
)

/*
*  Byte Slice
 */

func main() {
	f, err := os.Open("hellowworld.txt")

	if err != nil {
		fmt.Printf("%s \n", err)
		os.Exit(1)
	}

	defer f.Close()

	// byte slice
	b := make([]byte, 100)

	n, err := f.Read(b)

	fmt.Printf("%d : % x\n", n, b)
}
