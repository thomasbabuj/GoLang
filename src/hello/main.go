package main

import (
  "fmt"
)

// Using Defer keyword
// Passing unlimited nos of parameters to a function

func printer(format string, msgs ...string) {
  for _, msg := range msgs {
    fmt.Printf(format, msg)
  }
}

func main() {
  printer("%x \n", "Hello world!", "How are you?")
}

