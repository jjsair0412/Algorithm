package main

import "fmt"

func main() {
	var (
		H int
		M int
	)

	var tmp int = 24

	_, err := fmt.Scanf("%d", &H)
	if err != nil {
		fmt.Println("입력 에러 !")
		return
	}

	_, err = fmt.Scanf("%d", &M)
	if err != nil {
		fmt.Println("입력 에러 !")
		return
	}

	if H >= 0 && M >= 0 {
		if H <= 23 && M <= 59 {

		}
	} else {
		fmt.Println("잘못된 값을 입력하셨습니다.")
	}

}
