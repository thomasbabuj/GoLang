package main

import (
  "fmt"
)

// Since its a constant these variable can be only used inside main function
// If we want to use outside main we need to constant variable name first letter should be "Captital"
const (

  message = "The answer to life is %d \n"

  answer = 42
)

func main() {
  answer += 1
  fmt.Printf(message, answer)
}
