package main

import "fmt"

func main() {
	var (
		count  int
		result int
	)

	_, err := fmt.Scanf("%d", &count)
	if err != nil {
		fmt.Println("입력 에러")
		return
	}
	for i := 1; i <= count; i++ {
		result += i
	}

	fmt.Println(result)

}
