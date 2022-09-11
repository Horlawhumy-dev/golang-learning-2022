package main

import "fmt"

func main() {
	var s = []string{"Mon", "Tue", "Wed"}
	var y = make([]int, 0, 1)
	fmt.Printf("%v\t%d\t", s, len(s))
	fmt.Printf("%v\t%d\t", y, len(y))
	var arr = [2]string{"Arafat", "Olayiwola"}
	var arr2 = make([]string, len(arr), 4)
	//copy(arr2, arr)
	arr2[0] = "Arafat"
	arr2[1] = "Ola"
	//arr2[2] = "Olawumi"
	//var arr3 = arr2[2:4]
	var arr3 = make([]string, cap(arr2), cap(arr2)) //make([]Type, length, capacity)
	arr3[0] = "Olawumi"
	arr3[1] = "mynames"
	arr3[2] = "order"
	arr3[3] = "last"
	for _, e := range arr2 {
		fmt.Printf("\n%v\n", e)
	}
	fmt.Println(arr3)
}
