package main

import (
  "fmt"
)

/*
To assign automatic values to variables using iota keyword

using expression with iota
 */
const (

  message = "The answers are %d %d \n"

  answer1 = iota * 2

  answer2
)

func main() {

  fmt.Printf(message, answer1, answer2)
}
