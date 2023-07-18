package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		A     int
		count int
	)

	reader := bufio.NewReader(os.Stdin)

	slice := make([]int, 10)
	resultS := make([]int, 10)

	for i := range slice {
		fmt.Fscanf(reader, "%d\n", &A)
		slice[i] = A
	}

	for i, num := range slice {
		resultS[i] = num % 42
	}

	arr := buble(resultS)

	// fmt.Printf("%v\n", arr)
	for i := range arr {
		if i+1 == len(arr) {
			break
		}
		if arr[i] != arr[i+1] {
			count += 1
		}
	}

	fmt.Println(count + 1)

}

func buble(arr []int) []int {
	var count int
	for {
		if count != len(arr) {
			for i := range arr {
				if i+1 == len(arr) {
					break
				}
				if arr[i] > arr[i+1] {
					tmp := arr[i]
					arr[i] = arr[i+1]
					arr[i+1] = tmp
				}
			}
		} else {
			break
		}
		count += 1
	}
	return arr
}
