package main

import (
  "fmt"
)

type person struct{ //underlying type is a struct, you can create values of this type person now
  first string
  last string
  age int
}

type secretAgent struct {
  person
  first string
  ltk bool
}

func main() {

  sa1 := secretAgent{
    person: person{
      first: "James",
      last: "Bond",
      age: 32,
    },
    first: "something coll",
    ltk: true,
  }

  p2 := person{
    first: "Miss",
    last: "Moneypenny",
    age: 27,
  }

  fmt.Println(sa1) //create values of type person with fields in them
  fmt.Println(p2)
 
  fmt.Println(sa1.first, sa1.person.first, sa1.last, sa1.age)
  fmt.Println(p2.first, p2.last, p2.age)

}
