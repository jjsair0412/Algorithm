package main

import "fmt"

func main() {
	var (
		count int
	)

	_, err := fmt.Scanf("%d", &count)
	if err != nil {
		fmt.Println("입력 에러")
		return
	}

	switch count >= 1 && count <= 9 {
	case true:
		for i := 1; i <= 9; i++ {
			fmt.Printf("%d * %d = %d\n", count, i, i*count)
		}
	case false:
		fmt.Println("1에서 9 사이의 값을 입력해주세요..")
	}
}
