package main

import "fmt"

func main() {
	var (
		A int
		B int
		C int
	)

	fmt.Scanf("%d %d %d", &A, &B, &C)

	arr1 := [3]int{A, B, C}

	slice1 := make([]int, 3, 3)

	for i := 0; i <= 2; i += 1 {
		slice1[i] = arr1[i]
	}

	fmt.Printf("%d", slice1[i])

}
