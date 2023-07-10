package main

import "fmt"

func main() {

	var (
		score int
	)

	fmt.Scanf("%d", &score)

	switch score >= 0 && score <= 100 {
	case true:
		if score >= 90 && score <= 100 {
			fmt.Println("A")
		} else if score >= 80 && score <= 89 {
			fmt.Println("B")
		} else if score >= 70 && score <= 79 {
			fmt.Println("C")
		} else if score >= 60 && score <= 69 {
			fmt.Println("D")
		} else {
			fmt.Println("F")
		}
	case false:
		fmt.Println("잘못된 값을 입력하셨습니다.")
	}

}
