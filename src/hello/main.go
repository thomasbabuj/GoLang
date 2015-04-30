package main

import (
  "os"

)

// Using Defer keyword
//  defer section will be called after the return statement in a function
//  example of using defer with open and close a file
//  defer is how we deal with clean up in a function
 func printer(msg string) error {
  msgconvert := []byte(msg)
  f, err := os.Create("hellowworld.txt")
  if err != nil {
    return err
  }
  defer f.Close()

  f.Write(msgconvert)
  return err
}

func main() {

  printer("Hello, World!")
}
