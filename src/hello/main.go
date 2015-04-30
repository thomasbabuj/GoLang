package main

import (
  "fmt"
)

// Array , Slices
// Built-in types

/*
*  When we pass an array to a function, actually passes the copy of the array
*  so modifying the array in function will not change the value of the original array
*
 */

func printer(words [4]string) {
  for _, word := range words {
    fmt.Printf("%s ", word)
  }

  fmt.Printf("\n")

  words[2] = "blue"

}


func main() {
    // specify array size
    words := [4]string{"the", "quick", "brown", "fox"}

    printer(words)

    // Even though we change the value brown to blue in a printer function, now its shows the original array
    fmt.Printf("checking the words: %s", words[2])

}



