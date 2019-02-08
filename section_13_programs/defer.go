package main

import (
  "fmt"
)

func main() {
  defer foo()
  bar()
} // after exiting func main() the defered actions will be run -- ensures that the program does not waste resources and memory. Ex: a program opening files and never closing them. So right when you open a file you can have a function to close the file and defer it after the file was written.

func foo() {
  fmt.Println("foo")
}

func bar() {
  fmt.Println("bar")
}
