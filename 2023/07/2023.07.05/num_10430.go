package main

import "fmt"

func main() {
	var (
		A   int
		B   int
		C   int
		tmp int
	)

	fmt.Scanf("%d %d %d", &A, &B, &C)

	arr := [...]int{A, B, C}

	switch 2 <= A && C <= 10000 {
	case true:
		if arr[0] > arr[1] {
			tmp = arr[0]
			arr[0] = arr[1]
			arr[1] = tmp
		}
		if arr[1] > arr[2] {
			tmp = arr[1]
			arr[1] = arr[2]
			arr[2] = tmp
		}
		if arr[0] > arr[1] {
			tmp = arr[0]
			arr[0] = arr[1]
			arr[1] = tmp
		}
		fmt.Println(arr)

		// (A+B)%C
		fmt.Println((arr[0] + arr[1]) % arr[2])
		// ((A%C) + (B%C))%C
		fmt.Println(((arr[0] % arr[2]) + (arr[1] % arr[2])) % arr[2])
		// (A×B)%C
		fmt.Println((arr[0] * arr[1]) % arr[2])
		// ((A%C) × (B%C))%C
		fmt.Println(((arr[0] % arr[2]) * (arr[1] % arr[2])) % arr[2])

	case false:
		fmt.Println("잘못된 값을 입력하셨습니다. 조건 : 2 <= A && C <= 10000")
	}
}
