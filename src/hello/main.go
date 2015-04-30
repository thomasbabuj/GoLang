package main

import (
  "fmt"
)

// Using Defer keyword
//  defer section will be called after the return statement in a function
 func printer(msg string) error {
   // here we are adding new line after the function returns value
   // we can also add multiple defer. but it will stack up and follow a certain order when it runs
   defer fmt.Printf("\n")
   defer fmt.Printf("Add more \n")
  _, err := fmt.Printf("%s", msg)

  return err
}

func main() {

  printer("Hello, World!")
}
