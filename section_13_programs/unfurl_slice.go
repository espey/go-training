package main

import (
  "fmt"
)

func main() {
  foo() //every time you call a function you have to specify args, even if none
  xi := []int{2,3,4,5,6,7,8,9}
  z := "erin"
  x := sum(z, xi...) //unfurling the slice of ints
  fmt.Println("The total is", x)
  y := sum(z)
  fmt.Println("The total is", y) //you can pass nothing into the function as well, when asking for a variadic parameter you can pass in no argumet. Variadic means 0 or more of type int
  zz := sum(z, 4, 5, 6)
  fmt.Println("The total is", zz)
}

func foo() {
  fmt.Println("Hello, playground")
}

func sum(s string, x ...int) int { //unlimited number of ints, will store them as a slice of values
  fmt.Println(x)
  fmt.Printf("%T\n", x)
  fmt.Println(len(x))
  fmt.Println(cap(x))
  
  sum := 0
  for i, v := range x {
    sum += v
    fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
  }
  return sum
}
