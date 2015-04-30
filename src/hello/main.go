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
   // creating a slice with initial value zero and maximum 4
   // using copy() we can make copy of the slice

   words := make([]string, 0, 4)

   fmt.Printf("Size : %d , Capacity : %d \n", len(words), cap(words))

   words = append(words, "The")
   words = append(words, "Quick")
   words = append(words, "Brown")
   words = append(words, "Fox")

   newwords := make([]string, 4)
   copy(newwords, words)

   printer(newwords)
   newwords[2] = "Blue"

   printer(newwords)
   printer(words)

}


