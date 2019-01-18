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
  
  fmt.Println(p1.first_name)
  fmt.Println(p1.last_name) 
  for i, v := range p1.favFlavors {
    fmt.Println(i, v)
  }

  fmt.Println(p2.first_name)
  fmt.Println(p2.last_name) 
  for i, v := range p2.favFlavors {
    fmt.Println(i, v)
  }
}
