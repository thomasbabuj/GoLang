package main

import (
  "fmt"
)

// Using Defer keyword
// Declaring a variable called error "e"
// implicitly return values with out specify variable name
func printer(msg string) (e error) {

  // since we are getting value e here, this function automatically returns it.
  _, e = fmt.Printf("%s\n", msg)
  return
}

func main() {
  printer("Hello, World!")
}
