package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		N     int
		S     int
		input int
		count int
	)
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanf(reader, "%d\n", &N)

	slice := make([]int, N)

	for i := range slice {
		fmt.Fscanf(reader, "%d", &input)
		slice[i] = input
	}

	fmt.Fscanf(reader, "%d\n", &S)

	// 입력값 : 가동범위
	// 1 : 2
	// 2 : 4
	// 3 : 6
	// 4 : 8
	// 9 : 10

	for i := range slice {
		if i+1 == len(slice) {
			break
		}
		
		if  {
			break
		}

		if slice[i] < slice[i+1] {
			tmp := slice[i]
			slice[i] = slice[i+1]
			slice[i+1] = tmp
			count += 2
		}
	}

	for _, num := range slice {
		fmt.Printf("%d ", num)
	}
}
