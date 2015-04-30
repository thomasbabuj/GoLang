package main

import (
  "fmt"
)

//  multiple return values
 func printer(msg string) (string, error) {

  msg += "\n"

  _, err := fmt.Printf("%s \n", msg)

  return msg, err
}

func main() {

  appendedMsg ,err := printer("Hello, World");
  // %q form prints out a user friendly ( readable way)
  if err == nil {
    fmt.Printf("%q \n", appendedMsg)
  }

}
