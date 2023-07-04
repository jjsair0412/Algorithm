package main

import "fmt"

func main() {
	var (
		first  float64 //0xc0000160a8
		second float64
		// result float64
	)

	fmt.Scanf("%f %f", &first, &second)

	result := first / second
	// fmt.Printf("%.9f", first/second)
	fmt.Println(result)
}
