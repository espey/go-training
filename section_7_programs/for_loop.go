package main

import (
  "fmt"
)

func main() {
  x := 1
  for x < 10 { // x < 10 must evaluate to true; similar to a while loop
    fmt.Println(x)
    x++
  }
  fmt.Println("done.")
}
