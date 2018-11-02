package main

import "fmt"

func main() {
  fmt.Println("Hello everyone, this is the best yo")
  foo()
  fmt.Println("something more")

  for i := 0; i < 100; i++ {
    if i %2 == 0{
      fmt.Println(i)
    }
  }
  bar() 
}

func foo() {
  fmt.Println("I'm in foo")
}

func bar() {
  fmt.Println("Exiting yo")
}
