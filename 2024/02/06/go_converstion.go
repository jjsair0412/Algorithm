package main

import "fmt"

func main() {
	y := 42
	z := 42.0

	fmt.Printf("%v of type %T\n ", y, y)
	fmt.Printf("%v of type %T\n ", z, z)

	var m float32 = 43.742
	fmt.Printf("%v of type %T\n ", m, m)

	// z = m is'nt work

	z = float64(m)
	fmt.Printf("%v of type %T\n ", z, z)

}
