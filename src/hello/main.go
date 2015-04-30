package main

import (
  "fmt"
)

// Return a value
func printer(msg string) error {
  _, err := fmt.Printf("%s \n", msg)

  return err
}

func main() {

  printer("Hello, World");

}
