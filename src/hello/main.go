package main

import (
  "fmt"
)

// Array , Slices
// Built-in types

func main() {
    // since we are using "..." so we can store more values
    // there is no specific array size in this case
    words := [...]string{"the", "quick", "brown", "fox"}
    fmt.Printf("%s \n", words[2])
}



