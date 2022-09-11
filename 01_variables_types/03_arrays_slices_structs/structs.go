package main

import "fmt"

type FirstName struct {
	First string
}
type Profile struct {
	Name      string
	Age       int
	Height    float64
	IsMarried bool
}

func main() {
	arafat := newProfile(
		"Arafat",
		25,
		6.2,
		true)
	//ola := newProfile(
	//	"Olayiwola",
	//	20,
	//	arafat.changeHeight(),
	//	false)
	//olawumi := newProfile(
	//	"olawumi",
	//	20,
	//	changeHght(ola),
	//	false)
	person := Profile{"Arafat", 26, 6.5, false}
	fmt.Println(person)
	fmt.Println(arafat.Height)
	//fmt.Println(ola.Height)
	//fmt.Println(olawumi.height)
}

// constructor
func newProfile(name string, age int, height float64, isMarried bool) Profile {
	return Profile{
		name,
		age,
		height,
		isMarried,
	}
}

// method of struct
func (p Profile) changeHeight() float64 {
	return p.Height * 2
}

// function
func changeHght(p Profile) float64 {
	return p.Height * 2
}
