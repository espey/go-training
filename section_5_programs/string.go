package main

import (
  "fmt"
)

func main() {
  s := "Hello Playground"
  fmt.Println(s)
  fmt.Printf("%T\n", s)

  bs := []byte(s) //byte is an alais of type uint8
  fmt.Println(bs)
  fmt.Printf("%T\n", bs) //will print the ASCII letters as numbers: code scheme

  for i := 0; i < len(s); i++ {
    fmt.Printf("%#U ", s[i])
  }

  fmt.Println("")

  for i, v := range s {
    fmt.Println(i, v)
  }

  for i, v := range s {
    fmt.Printf("at index position %d we have hex %#x\n", i, v)
  }
}
