package main

import (
  "fmt"
)

func main() {

  m := struct{
    colors []string
    friends map[string]int
  }{
    colors: []string{"red", "green", "blue"},
    friends: map[string]int{
      "kelly": 2,
      "dave": 4,
    },
  }

  for _, v := range m.colors {
    fmt.Println(v)
  }

  for k, v := range m.friends {
    fmt.Println(k, v)
  }

}
