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
		A -= 1
		B -= 1
		for i := A; i < B; i++ {
			for j := B; j > i; j-- {
				tmp := slice[j-1]
				slice[j-1] = slice[j]
				slice[j] = tmp
			}
		}
	}

	for i := range slice {
		fmt.Printf("%d ", slice[i])
	}
}
