package main

import "fmt"

func main() {
	var (
		first  float64
		second float64
	)

	fmt.Scanf("%f %f", &first, &second)

	result := first / second
	fmt.Println(result)
}
