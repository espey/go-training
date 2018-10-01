package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, I'm learning go programming")
	foo()
	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("Hello there", 42, true)

	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("and then we exited")
}
