package main

import "fmt"

func main() {
	var N int

	_, err := fmt.Scanf("%d", &N)
	if err != nil {
		fmt.Println("입력 에러 발생")
		return
	}

	for i := 1; i <= N; i++ {
		if i%4 == 0 {
			fmt.Println("long")
		}
		if i == N {
			fmt.Println("int")
		}
	}

}
