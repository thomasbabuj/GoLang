package main

import (
  "fmt"
)

func main() {
  // Go String
  /* declaring string with backquotes '`' means, anything in between start and end backquotes will be treated as a full string. So we dont need to worry about escaping the strings for single or double quotes etc ..
  */

  atoz := `the quick brown fox jumps over the lazy dog \n`

  // This example \n wont be removed since we declare the string with '`'
  fmt.Printf("String : %s \n", atoz)
}
