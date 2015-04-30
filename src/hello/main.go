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

/*
Currently we already know the value of the slice ( static slice )
 */

func printer(words []string) {
   for _,word := range words {
     fmt.Printf(" %s ", word)
   }

   fmt.Printf("\n")
}

func main() {
   // createing a new slice with using make
   // created a new slice with the max size of 4
   words := make([]string, 4)
   words[0] = "The"
   words[1] = "Quick"
   words[2] = "Brown"
   words[3] = "Fox"

   printer(words)
}


