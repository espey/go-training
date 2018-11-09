package main

import (
  "fmt"
)

var a int = 42
var b int = a << 1


func main() {

fmt.Printf("%d\t%b\t%x", a, a, a)
fmt.Println("")
fmt.Printf("%d\t%b\t%x", b, b, b)

 
}
