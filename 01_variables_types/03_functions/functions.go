package main

import "fmt"

func main() {
	res := add(3, 2)
	fmt.Println(res)

	anonymous := func(x int, y int) int {
		return x - y
	}

	fmt.Println(anonymous(3, 2))

	func(x string, y string) {
		fmt.Println(x + y)
		return
	}("3", "2")
	x := 3 // positive int value
	fmt.Println(x)
	negate(&x) //negative int value
	fmt.Println(x)
}

func add(x int, y int) int {
	return x + y
}

func negate(x *int) {
	neg := -*x // negating whatever value got from dereferencing pointer x
	*x = neg   // storing negative value to pointer x
}
