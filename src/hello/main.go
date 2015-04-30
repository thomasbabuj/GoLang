package main

import (
  "fmt"
)

// Using Defer keyword
// Passing unlimited nos of parameters to a function

func printer(msgs ...string) {
  for _,msgs := range msgs {
    fmt.Printf("%s \n", msgs)
  }
}

func main() {
  printer("Hello, world!", "How are u doing")
}

