package main

import "fmt"

func main() {

	var (
		input int
	)

	fmt.Scanf("%d", &input)

	if input >= 1 && input <= 4000 {
		if input%4 == 0 && input%100 != 0 {
			fmt.Println(1)
		} else if input%400 == 0 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	} else {
		fmt.Println("잘못된 값을 입력하셨습니다.")
	}
}
