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
   // but using append function we can add more values
   //
   words := make([]string, 0, 4)
   words = append(words, "The")
   words = append(words, "Quick")
   words = append(words, "Brown")
   words = append(words, "Fox")
   words = append(words, "bbbbb")
   words = append(words, "miououou")

   printer(words)

   fmt.Printf("Size : %d", len(words))
}


