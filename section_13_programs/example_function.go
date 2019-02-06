package main

import(
  "fmt"
)

func main() {
  foo()
  bar("Erin")
  s1 := woo("Moneypenny")
  fmt.Println(s1)
  x, y := mouse("Ian", "Flemming")
  fmt.Println(x)
  fmt.Println(y)
}

// func (r reciever) identifier(parameters) (return(s)) {.....}

func foo() { //basic func
  fmt.Println("hello from foo")
}

// everything in go is passed by value
func bar(s string) { //defined with string as a paramater, takes in arg
  fmt.Println("Hello", s)
}

func woo(st string) string { //return func
  return fmt.Sprint("Hello from woo, ", st)
}

func mouse(fn, ln string) (string, bool) {
  a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
  b := false
  return a, b
}
