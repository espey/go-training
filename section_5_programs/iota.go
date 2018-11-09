package main

import (
  "fmt"
)

const (
  a = iota // 0
  b = iota // 1 iota will increment by one
  c = iota // 2
)

const (
  d = iota // 0
  e = iota // 1 iota will reset when you define a block of constants
  f = iota // 2

)

func main() {

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
  fmt.Println(e)
  fmt.Println(f)

  fmt.Printf("%T\n", a)
  fmt.Printf("%T\n", b)
  fmt.Printf("%T\n", c)
}
