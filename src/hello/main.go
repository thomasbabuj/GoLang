package main

import (
  "fmt"
)

func main() {
  // Go String
  atoz := "the quick brown fox jumps over the lazy dog \n"

  // In Go we can use "_" as variable name to specify that we are only using one value.
  for _, r := range atoz {

    // Printing the index and chars
    // Since we use _ we are only getting the chars and we omit the index of the char
    fmt.Printf(" %c \n", r)
  }
}
