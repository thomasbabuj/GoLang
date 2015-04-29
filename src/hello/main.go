package main

import (
  "fmt"
)

func main() {
  // Go String
  /* declaring string with back quotes '`' means, anything in between start and end back quotes will be treated as a full string. So we don't need to worry about escaping the strings for single or double quotes etc ..
  */

  atoz := `the "quick" brown fox jumps over the lazy dog \n`

  // This example \n wont be removed since we declare the string with '`'
  fmt.Printf("String : %s \n", atoz)
}
