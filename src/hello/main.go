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
   // when the slice size over pass with the capacity go automatically increase the size
   words := make([]string, 0, 4)

    fmt.Printf("Size : %d , Capacity : %d \n", len(words), cap(words))

   words = append(words, "The")
   words = append(words, "Quick")
   words = append(words, "Brown")
   words = append(words, "Fox")

    fmt.Printf("Size : %d , Capacity : %d \n", len(words), cap(words))

   words = append(words, "Jumps")
   words = append(words, "Over")
   words = append(words, "The")
   words = append(words, "lazy")

   fmt.Printf("Size : %d , Capacity : %d \n", len(words), cap(words))

    words = append(words, "Dog")

    fmt.Printf("Size : %d , Capacity : %d \n", len(words), cap(words))

    printer(words)
}


