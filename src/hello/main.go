package main

import (
  "fmt"
  "os"
)

func main() {
  // Go Printf returns also returns errors
  if  numberBytes, theError := fmt.Printf("Hello, world! \n"); theError != nil {
    os.Exit(1)
  } else {
     // since numberBytes was declared in if scope so its not available here
    fmt.Printf("Printed %d bytes \n", numberBytes)
  }



}
