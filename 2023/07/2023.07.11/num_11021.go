package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		count int
		A     int
		B     int
	)

	reader := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscanf(reader, "%d\n", &count)
	if err != nil {
		fmt.Println("입력 에러 발생")
		return
	}

	for i := 1; i <= count; i++ {
		_, err := fmt.Fscanln(reader, &A, &B)
		if err != nil {
			fmt.Println("for문 내부 입력에러 발생")
			fmt.Println(err)
			return
		}
		result := A + B
		_, err = fmt.Fprintf(os.Stdout, "Case #%d: %d\n", i, result)
		if err != nil {
			fmt.Println("for문 내부 출력에러 발생")
			fmt.Println(err)
			return
		}

	}
}
