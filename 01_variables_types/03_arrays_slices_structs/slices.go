package main

import "fmt"

func main() {
	var s = []string{"Mon", "Tue", "Wed"}
	var y = make([]int, 0, 1)
	fmt.Printf("%v\t%d\t", s, len(s))
	fmt.Printf("%v\t%d\t", y, len(y))
}
