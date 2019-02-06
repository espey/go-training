package main

import (
  "fmt"
)

func main() {
  foo() //every time you call a function you have to specify args, even if none
  x := sum(2,3,4,5,6,7,8,9)
  fmt.Println("The total is", x)
}

func foo() {
  fmt.Println("Hello, playground")
}

func sum(x ...int) int { //unlimited number of ints, will store them as a slice of values
  fmt.Println(x)
  fmt.Printf("%T\n", x)
  
  sum := 0
  for i, v := range x {
    sum += v
    fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
  }
  return sum
}
