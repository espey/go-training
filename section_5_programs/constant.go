package main

import (
  "fmt"
)

const a = 42
const b = 42.78
const c = "James Bond"

const (
  d int = 1
  e = 2
)

func main() {
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
  fmt.Println(e)

  fmt.Printf("%T\n", a)
  fmt.Printf("%T\n", b)
  fmt.Printf("%T\n", c)
  fmt.Printf("%T\n", d)
  fmt.Printf("%T\n", e)
}
