package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var (
		A        int
		max      int
		location int
	)
	slice := make([]int, 9)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i <= 8; i++ {
		_, err := fmt.Fscanf(reader, "%d\n", &A)
		if err != nil {
			fmt.Println("입력 에러")
			fmt.Println(err)
			return
		}
		slice[i] = A
	}

	for i, realNum := range slice {
		if i == 0 {
			max = slice[0]
		} else {
			if max < realNum {
				max = realNum
			}
		}

	}

	for i, realNum := range slice {
		if realNum == max {
			location = i + 1
		}
	}

	fmt.Println(max)
	fmt.Println(location)

}
