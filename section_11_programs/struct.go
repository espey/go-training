package main

import (
  "fmt"
)

type person struct{ //underlying type is a struct, you can create values of this type person now
  first string
  last string
  age int
}

func main() {

  p1 := person{
    first: "James",
    last: "Bond",
    age: 32,
  }

  p2 := person{
    first: "Miss",
    last: "Moneypenny",
    age: 27,
  }

  fmt.Println(p1) //create values of type person with fields in them
  fmt.Println(p2)
 
  fmt.Println(p1.first, p1.last, p1.age)
  fmt.Println(p2.first, p2.last, p2.age)

}
