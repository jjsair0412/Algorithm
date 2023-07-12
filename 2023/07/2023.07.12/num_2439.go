package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var count int

	reader := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscanf(reader, "%d", &count)
	if err != nil {
		fmt.Println("입력 에러발생")
		return
	}

	switch count <= 100 && count >= 1 {
	case true:
		for i := count; i > 0; i-- {
			for j := 1; j <= i-1; j++ {
				fmt.Print(" ")
			}
			for t := count; t >= i; t-- {
				fmt.Print("*")
			}
			fmt.Println()
		}
	case false:
		fmt.Println("1부터 100 사이의 값을 넣어주세요.")
	}
}
