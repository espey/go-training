package main

import (
  "fmt"
)

func main() {
  switch {
  case false:
    fmt.Println("this should not print")
  case (2 == 4):
    fmt.Println("this should not print")
  default: //this is what happens if nothing else evaluates to true
    fmt.Println("this is default")
  }
}
