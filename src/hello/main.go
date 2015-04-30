package main

import (
  "fmt"
)

// since both two parameters are using same type
// we can use the following syntax
func printer(msg , msg2 string ,repeat int) {
  for repeat > 0 {
    fmt.Printf("%s", msg)
    fmt.Printf(" %s", msg2)
    fmt.Printf("\n")
    repeat -= 1
  }
}

func main() {

  // Defining functions in Go
  printer("Hello", "world", 5)
}
