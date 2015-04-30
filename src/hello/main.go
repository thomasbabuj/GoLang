package main

import (
  "fmt"
  "os"
)

func main() {

    n, err := fmt.Printf("Hello, World! \n")

    switch {
      case err != nil :
        os.Exit(1)
      case n == 0:
        fmt.Printf("No bytes output \n")
      case n != 4 :
        fmt.Printf("Wrong number of characters \n")
      default:
        fmt.Printf("Ok!!! \n")
    }

    fmt.Printf("\n")
}
