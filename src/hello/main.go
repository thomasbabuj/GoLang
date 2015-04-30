package main

import (
  "fmt"
)

// Using Defer keyword
// Declaring a variable called error "e"
// implicitly return values with out specify variable name
func printer(msg string) (returnedMessage string, e error) {

  // since we are getting value e here, this function automatically returns it.
  _, e = fmt.Printf("%s\n", msg)
  returnedMessage = "Yes I did it"
  return
}

func main() {
 value1,value2 := printer("Hello, World!")

  if (value2 == nil ) {
    fmt.Printf("I am in if \n")
    fmt.Printf(" %s \n", value1 )
  } else {
    fmt.Printf("I am in else \n")
    fmt.Printf(" %s %s", value1, value2 )
  }
}
