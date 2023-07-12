package main

import (
	"bufio"
	"fmt"
	"io"
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
			if err == io.EOF {
				break
			} else {
				fmt.Println("입력 에러발생")
				fmt.Println(err)
				return
			}
		}

		switch A > 0 && B < 10 {
		case true:
			_, err = fmt.Fprintln(os.Stdout, A+B)
			if err != nil {
				fmt.Println("출력 에러발생")
				fmt.Println(err)
				return
			}
		case false:
			fmt.Println("올바른 값을 입력해주세요 . A는 0보다 크고 B는 10보다 작아야 합니다.")

		}

	}
}
