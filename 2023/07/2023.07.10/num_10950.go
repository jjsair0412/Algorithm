package main

import "fmt"

func main() {
	var (
		count int
		A     int
		B     int
	)

	_, err := fmt.Scanf("%d\n", &count)
	if err != nil {
		fmt.Println("입력 에러 발생")
		return
	}

	for i := 0; i < count; i++ {
		_, err := fmt.Scanf("%d %d\n", &A, &B)
		if err != nil {
			fmt.Println("입력 에러 발생")
			return
		}

		fmt.Printf("%d\n", A+B)
	}
}
