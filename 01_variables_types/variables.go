package main

import "fmt"

var z bool // declaration
const (
	i1 = iota << 10
	i2 = iota << 10
	i3 = iota << 10
	i4 = iota << 10
)

func main() {
	x := 43      // short declaration
	y := "Hello" // short declaration
	z = true     // initialization
	m := false
	n := z && m || z
	var p float32 = 9.00
	q := 86.0393 //float64
	var i int = 83333344445
	var c1 complex64 = complex(10, 3) // complex64(both im & re are float32), complex128(float64)
	fmt.Println(n, i, p, c1, q)
	const b = "Hello"
	const a = b // two constants can be assigned to each other
	fmt.Println(x, y, z)

	b1 := 11
	b2 := b1 << 10 // bit shifting

	fmt.Printf("%d\t\t\t%b\n", b1, b1)
	fmt.Printf("%d\t\t\t%b\n", b2, b2)
	fmt.Println(i1, i2, i3, i4)
	fmt.Printf("%b\n", i3)
	fmt.Printf("%b\n", i4)
}
