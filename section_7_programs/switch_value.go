package main

import (
  "fmt"
)

func main() {
  switch "Bond" {
  case "Moneypenny":
    fmt.Println("miss money")
  case "Bond":
    fmt.Println("bond james")
  case "q":
    fmt.Println("this is q")
  default: //this is what happens if nothing else evaluates to true
    fmt.Println("this is default")
  }
}
