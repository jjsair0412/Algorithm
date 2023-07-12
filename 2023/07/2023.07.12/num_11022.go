package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		A     int
		B     int
		count int
	)

	reader := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscanf(reader, "%d\n", &count)
	if err != nil {
		fmt.Println("초기 값 입력 에러발생")
		fmt.Println(err)
		return
	}

	for i := 1; i <= count; i++ {
		_, Rerr := fmt.Fscanln(reader, &A, &B)
		if Rerr != nil {
			fmt.Println("입력 에러발생")
			fmt.Println(Rerr)
			return
		}
		result := A + B
		_, Werr := fmt.Fprintf(os.Stdout, "Case #%d: %d + %d = %d\n", i, A, B, result)
		if Werr != nil {
			fmt.Println("출력 에러발생")
			fmt.Println(Werr)
			return
		}
	}
}
