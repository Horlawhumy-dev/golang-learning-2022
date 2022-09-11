package main

import "fmt"

var y string

func main() {
	n, _ := fmt.Println("Hello World")
	fmt.Println(n)
	age := 27
	ageptr := &age
	fmt.Println(age, ageptr)
	*ageptr = 28
	fmt.Println(age)
	x := 64
	fmt.Println(x)
	x = 4
	foo()
	fmt.Println(y)

	ages := 20
	increment(&ages)
	fmt.Println(ages)
	//age is now 21

}
func increment(a *int) {
	*a = *a + 1
}

func foo() {
	y = "Hello"
}
