package main

import "fmt"

func main() {
	if x := 4; x > 5 {
		fmt.Println("x is greater than  5")
	} else if x < 5 {
		fmt.Println("x is less than 5")
	} else {
		fmt.Println("x is equal to 5")
	}
}
