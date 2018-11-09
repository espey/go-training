package main

import (
  "fmt"
)

const (
  _ = iota
  ep = 1 << (iota * 10)
  dp = 1 << (iota * 10)
  kp = 1 << (iota * 10)
)

func main() {
  x := 4
  fmt.Printf("%d\t\t%b", x, x)

  fmt.Println("")

  y := x << 1 // shift the bit over from 100 to 1000
  fmt.Printf("%d\t\t%b", y, y)

  kb := 1024
  mb := 1024 * kb
  gb := 1024 * mb

  fmt.Println("")

  fmt.Printf("%d\t\t\t%b\n", kb, kb)
  fmt.Printf("%d\t\t\t%b\n", mb, mb)
  fmt.Printf("%d\t\t%b\n", gb, gb)

  fmt.Printf("%d\t\t\t%b\n", ep, ep)
  fmt.Printf("%d\t\t\t%b\n", dp, dp)
  fmt.Printf("%d\t\t%b\n", kp, kp)

}
