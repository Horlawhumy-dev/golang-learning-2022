package main

import "fmt"

var a string = "Arafat"

func main() {
	var ptr *string = &a // storing the memory address of variable "a" to the pointer ptr
	var ptr2 *string     // return nil by default
	value := *ptr        // returns value from a memory address
	fmt.Printf("a=%v pointer=%v value=%v\n", a, ptr, value)
	fmt.Println(&a)
	fmt.Println(ptr2)

}
