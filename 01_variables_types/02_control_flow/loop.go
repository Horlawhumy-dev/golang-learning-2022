package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		for j := i + 1; j < 5; j++ {
			fmt.Printf("%#U\t%#x", i, i)
		}
		fmt.Println(i)
	}
}
