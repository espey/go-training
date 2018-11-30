package main

import (
  "fmt"
)

func main() {

  x := "James Bond"
  if x == "James Bond" {
    fmt.Println(x)
  } else if x == "James Bland" {
    fmt.Println("Not James Bond")
  } else {
    fmt.Println("No Spys here!")
  }
}
