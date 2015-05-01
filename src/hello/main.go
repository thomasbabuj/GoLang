package main

import (
	"fmt"
	"os"
)

/*
*  Byte Slice
 */

func main() {
	f, err := os.Open("test.txt")

	if err != nil {
		fmt.Printf("%s \n", err)
		os.Exit(1)
	}

	defer f.Close()

	// byte slice
	b := make([]byte, 100)

	n, err := f.Read(b)

	// we can't print this as string
	fmt.Printf("%d : % c\n", n, b)

	// we need to create a new string variable
	stringVaribale := string(b)

	fmt.Printf("%d %s \n", n, stringVaribale)

	// writing into to a file
	// since wirte function need byte type we need to convert the string
	// to byte slice
	someString := "foo bar"
	f.Write([]byte(someString))
}
