package main

import "fmt"

func main() {
	var (
		x int
		y int
	)

	_, err := fmt.Scanf("%d\n", &x)
	if err != nil {
		fmt.Println("입력 에러:", err)
		return
	}

	_, err = fmt.Scanf("%d\n", &y)
	if err != nil {
		fmt.Println("입력 에러:", err)
		return
	}

	if x > 0 && y > 0 {
		fmt.Println("1")
	} else if x < 0 && y > 0 {
		fmt.Println("2")
	} else if x < 0 && y < 0 {
		fmt.Println("3")
	} else if x > 0 && y < 0 {
		fmt.Println("4")
	}

}
