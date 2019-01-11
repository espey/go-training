package main

import (
  "fmt"
)

func main() {
  m := map[string]int{
    "James":32,
    "Miss Moneypenny":27,
  }

  fmt.Println(m)
  fmt.Println(m["James"])
  fmt.Println(m["Barnabas"])
  v, ok := m["Barnabas"]
  fmt.Println(v)
  fmt.Println(ok)

  m["todd"] = 33 //adding a value to the map
 
  if v, ok := m["Barnabas"]; ok {
    fmt.Println("success, exists", v)
  }

  if v, ok := m["Miss Moneypenny"]; ok {
    fmt.Println("success, exists", v)
  }

  for k, v := range m { // printing out the whole map
    fmt.Println(k, v)
  }
  
  delete(m, "James")
  fmt.Println(m) //removed James from the map

  delete(m, "Ian Fleming")
  fmt.Println(m) //you can delete entries from a map that are not even in the map


  if v, ok := m["Miss Moneypenny"]; ok { //verify you are deleting something
    fmt.Println("value:", v)
    delete(m, "Miss Moneypenny")
  }
  
  fmt.Println(m)
}
