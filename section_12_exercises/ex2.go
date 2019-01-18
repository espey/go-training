package main

import (
  "fmt"
)

type person struct{
  first_name string
  last_name string
  favFlavors []string
}


func main() {
  p1 := person{
    first_name: "John",
    last_name: "Doe",
    favFlavors: []string{
                  "vanilla", 
                  "chocolate",
                },
  }


   p2 := person{
    first_name: "Mary",
    last_name: "Moore",
    favFlavors: []string{
                  "mint", 
                  "cherry",
                  "strawberry",
                },
  }

  m :=  map[string]person{
    p1.last_name:p1,
    p2.last_name:p2,
  }
  
  for _, v := range m {
    fmt.Println(v.first_name)
    fmt.Println(v.last_name)
    for i, n := range v.favFlavors{
      fmt.Println(i, n)
    }
    fmt.Println("============")
  }
}
