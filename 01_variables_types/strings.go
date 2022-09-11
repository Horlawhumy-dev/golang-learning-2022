package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = "A-r-a-f-a-t"
	fmt.Println(name)
	var name2 = name[0:4]
	//name = "Ola"
	fmt.Println(name2)
	fmt.Println(name)
	fmt.Println(strings.Count(name, strings.ToUpper("a")))
	//arr := strings.Split(name, "-")
	names := []string{"Arafat", "Olayiwola"}
	//names[2] = "Ola"
	namesOutput := strings.Join(names, " ")
	fmt.Println(namesOutput)
}
