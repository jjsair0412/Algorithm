package main

import "fmt"

func main() {

	var (
		X      int // 영수증에 적힌 총 금액
		Y      int // 영수증에 적힌 물건 종류의 수
		a      int // 물건 가격
		b      int // 물건 개수
		result int
	)

	_, err := fmt.Scanf("%d\n", &X)
	if err != nil {
		fmt.Println("입력 에러")
		return
	}

	_, err = fmt.Scanf("%d\n", &Y)
	if err != nil {
		fmt.Println("입력 에러")
		return
	}

	for i := 0; i < Y; i++ {
		_, err = fmt.Scanf("%d %d\n", &a, &b)
		if err != nil {
			fmt.Println("입력 에러")
			return
		}
		result += a * b
	}

	switch result == X {
	case true:
		fmt.Println("Yes")
	case false:
		fmt.Println("No")
	}
}
