package main

import "fmt"

func main() {
	var arr [7]int
	arr = [7]int{1, 2, 4, 5, 6, 7, 8}
	//var arr3 []int = arr
	arr2 := []int{1, 2, 4, 5, 6, 7, 8}
	for i := 0; i < len(arr2); i++ {
		if arr[i]%2 == 0 {
			fmt.Printf("%d is even\n", arr2[i])
		} else {
			fmt.Printf("%d is odd\n", arr2[i])
		}
	}
	fmt.Println(len(arr2))
	for i, e := range arr {
		fmt.Println(i, e)
	}
}
