package main

import (
  "fmt"
)

func main() {
  // Go String
  atoz := "the quick brown fox jumps over the lazy dog \n"

  for i, r := range atoz {
    // Printing the index and chars
    fmt.Printf("%d %c \n", i, r)
  }
}
