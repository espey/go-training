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

  m["erin"] = []string{"books", "music", "sleeping"}

  for k,v := range m {
    fmt.Println("This is the record for", k)
    fmt.Println("\t", v)
  }

  delete(m, "erin")
  fmt.Println("\n", "erin should be deleted from the map")

  for k,v := range m {
    fmt.Println("This is the record for", k)
    fmt.Println("\t", v)
  }
}

