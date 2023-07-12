package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		A int
		B int
	)
	reader := bufio.NewReader(os.Stdin)

	for {
		_, err := fmt.Fscanf(reader, "%d %d\n", &A, &B)
		if err != nil {
			fmt.Println("입력 에러")
			fmt.Print(err)
			return
		}

		if A == 0 && B == 0 {
			break
		} else {
			fmt.Fprintln(os.Stdout, A+B)

		}
	}
}
