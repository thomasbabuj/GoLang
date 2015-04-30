package main

import (
  "fmt"
)

func main() {

  // Looping in Go
  //  simultaneous variable assignment
  for i, j := 0, 1; i < 10; i, j = i+1, j*2 {
    fmt.Printf("%d Hello, World \n", j)
  }

}
