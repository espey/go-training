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
  case (3 == 3):
    fmt.Println("prints")
    fallthrough // the next statement will only run if this is here
  case (4 == 4):
    fmt.Println("also true, does it print?")
    fallthrough // if this is here, it will print the next even if false...
  case (7 == 9):
    fmt.Println("not true")
    fallthrough
  case (11 == 14):
    fmt.Println("also not true")
    fallthrough
  default: //this is what happens if nothing else evaluates to true
    fmt.Println("this is default")
  }
}
