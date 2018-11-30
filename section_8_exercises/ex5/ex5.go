package main

import (
  "fmt"
)

func main() {
  for i := 10; i <= 100; i++ {
    j := i%4
    fmt.Printf("When %v is divided by 4, the remainder (aka modulus) is %v\n", i, j)
  }
}
