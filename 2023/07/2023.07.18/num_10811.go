package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		N int
		M int
		A int
		B int
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	slice := make([]int, N)

	for i := range slice {
		slice[i] = i + 1
	}

	for h := 0; h < M; h++ {
		fmt.Fscanf(reader, "%d %d\n", &A, &B)
		for i := A; i < B-1; i++ {
			for j := A; j < i; j-- {
				if j == 0 {
					break
				}
				slice[i] = slice[j]
				slice[j] = slice[i]
			}
		}
	}

	for i := range slice {
		fmt.Printf("%d ", slice[i])
	}
}
