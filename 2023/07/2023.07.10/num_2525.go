package main

import "fmt"

func main() {
	var (
		H  int
		M  int
		tm int
	)

	_, err := fmt.Scanf("%d", &H)
	if err != nil {
		fmt.Println("입력 에러 !")
		return
	}

	_, err = fmt.Scanf("%d\n", &M)
	if err != nil {
		fmt.Println("입력 에러 !")
		return
	}

	_, err = fmt.Scanf("%d", &tm)
	if err != nil {
		fmt.Println("입력 에러 !")
		return
	}

	if H >= 0 && M >= 0 {
		if H <= 23 && M <= 59 {
			switch M+tm >= 60 {
			case true:
				if tm >= 60 {
					H += ((M + tm) / 60)
					M = (M + tm) - (60 * ((M + tm) / 60))
					if M == 60 {
						M = 0
					}
				} else {
					H += 1
					M = (M + tm) - 60
				}
			case false:
				M += tm
			}

			if M < 0 {
				M += 60
			}

			if H > 24 {
				H -= 24
			} else if H == 24 {
				H = 0
			}
			fmt.Printf("%d %d", H, M)
		}
	} else {
		fmt.Println("잘못된 값을 입력하셨습니다.")
	}

}
