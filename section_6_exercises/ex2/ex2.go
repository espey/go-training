package main

import (
  "fmt"
)

func main() {
  a := 12
  b := 14
  c := a == b //false
  d := a <= b //true
  e := a >= b //false
  f := a != b //true
  g := a < b //true
  h := a > b //false

  fmt.Println(c)
  fmt.Println(d)
  fmt.Println(e)
  fmt.Println(f)
  fmt.Println(g)
  fmt.Println(h)

}
