package main

import (
  "fmt"
)

func main() {
  if x := 42; x == 2 { // you can put two statements on one line with a ;
    fmt.Println("false") // wont print
  }

  // fmt.Println(x) -- this won't run because x is limited in scope to that if statement

}
