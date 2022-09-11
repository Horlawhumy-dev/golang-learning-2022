package main

import "fmt"

func main() {
	var nameDict = make(map[string]int)
	//nameDicts := map[string] int{"age": 29}
	nameDict["age"] = 27

	fmt.Println(nameDict) //map[age:27]
}
