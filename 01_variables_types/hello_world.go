package main

import "fmt"

var y string

func main() {
	n, _ := fmt.Println("Hello World")
	fmt.Println(n)

	x := 64
	fmt.Println(x)
	x = 4
	foo()
	fmt.Println(y)

}

func foo() {
	y = "Hello"
}
