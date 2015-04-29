package main

import "fmt"

// Using a constant variable
//
const (

  message = "The answer to life is %d \n"

  answer = 42
)

func main() {

  fmt.Printf(message, answer)
}
