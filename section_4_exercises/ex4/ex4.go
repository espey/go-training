package main

import (
  "fmt"
)

type unicorn int
var x unicorn

func main() {

  fmt.Println(x)
  fmt.Printf("%T\n", x)
  x = 42
  fmt.Println(x)

}
