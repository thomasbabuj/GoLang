package main

import (
  "fmt"
)

// Using Deffer keyword
//  defer section will be called after the return statement in a function
 func printer(msg string) error {
   // here we are adding new line after the function returns value
   defer fmt.Printf("\n")
  _, err := fmt.Printf("%s", msg)

  return err
}

func main() {

  printer("Hello, World!")
}
