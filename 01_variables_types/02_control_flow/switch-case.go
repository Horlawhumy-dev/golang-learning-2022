package main

import "fmt"

func main() {
	day := "Saturday"
	switch day {
	case "Monday":
		fmt.Println("Monday")
	case "Saturday":
		fmt.Println("Saturday")
	case "Tuesday":
		fmt.Println("Tuesday")
	default:
		fmt.Println("Sunday")
	}
}
