package main

import (
  "fmt"
)

func main() {
  m := map[string][]string{
    "bond_james": {"shaken, not stirred", "Martinis", "Women"},
    "moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
    "no_dr": {"Being evil", "Ice cream", "Sunsets"},
  }
  
  for k, v := range m {
    fmt.Println("This is the record for", k)
      for integer, info := range v {
        fmt.Println("\t", integer, info)
      }
  }
}
