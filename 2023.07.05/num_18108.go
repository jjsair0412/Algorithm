package main

import "fmt"

func main() {

	var (
		y int
		// resultB bool
	)
	fmt.Scanf("%d", &y)
	switch y >= 1000 && y <= 3000 {
	case true:
		fmt.Println(y - 543)
		// resultB = true
	case false:
		fmt.Println("잘못된 값을 입력하였습니다.")
		y = 0
	}

	// for {
	// 	fmt.Println("계산대상 년도를 입력해주세요,,")

	// 	if y == 0 {
	// 		fmt.Scanf("%d", &y)
	// 	}
	// 	fmt.Scanf("%d", &y)

	// 	switch y >= 1000 && y <= 3000 {
	// 	case true:
	// 		fmt.Println(y - 543)
	// 		resultB = true
	// 	case false:
	// 		fmt.Println("잘못된 값을 입력하였습니다.")
	// 		y = 0
	// 	}

	// 	if resultB {
	// 		break
	// 	}
	// }

}
