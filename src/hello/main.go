package main

import (
  "fmt"
)

// Array , Slices
// Built-in types

/*
*  Usually we use slice over array. Slice is a window into underlying array
*  In slice we dont specify the length for the array Eg:  words := []string{"kkk", "ddfsfd"}
*  Since we are using slice, we can change original array value in a function
*/

func printer(words []string) {
   for _,word := range words {
     fmt.Printf(" %s ", word)
   }

   fmt.Printf("\n")
}

func main() {
  words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}

  fmt.Printf("Array Size : %d \n", len(words))

  printer(words[2:4])
}


