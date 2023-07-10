package main

import "fmt"

func main() {
	var (
		H int
		M int
	)

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
			if M < 45 {
				H -= 1
				M = 60 - (45 - M)
			} else if M >= 45 {
				M -= 45
			}
			if H < 0 {
				H += 24
			}
			fmt.Printf("%d %d", H, M)
		}
	} else {
		fmt.Println("잘못된 값을 입력하셨습니다.")
	}

}
