package main

import (
  "fmt"
)

func main() {
  // Boolean
  // In Go if you declare a variable without assign any values then by default it gets "Zero values"
  var isTrue bool

  // Since we didn't assign any value for isTrue variable now we get false as our output
  fmt.Printf("Value : %t\n", isTrue)
}
