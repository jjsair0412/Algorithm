package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		N     int
		M     int
		count int
		a     int
		b     int
	)

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	_, err := fmt.Fscanf(reader, "%d %d\n", &N, &M)
	if err != nil {
		fmt.Println("N , M 입력에러 발생")
		fmt.Println(err)
		return
	}

	switch 1 <= N && 100 >= N && 1 <= M && 100 >= M {
	case true:

	case false:
		fmt.Println("N , M 입력 조건 : 1 <= N && 100 >= N && 1 <= M && 100 >= M ")
	}
	baguniNum := make([]int, N)
	for i := 0; i < N; i++ {
		baguniNum[i] = i + 1 // 바구니에 공 넣기
	}

	for {
		if count != M {
			var tmp int = 0
			_, err := fmt.Fscanf(reader, "%d %d\n", &a, &b)
			if err != nil {
				fmt.Println("a , b 입력에러 발생")
				fmt.Println(err)
				return
			}
			switch a >= 1 && a <= b && b <= N {
			case true:
				tmp = baguniNum[a-1]
				baguniNum[a-1] = baguniNum[b-1]
				baguniNum[b-1] = tmp
			case false:
				fmt.Println("a, b 입력조건 : a >= 1 && a <= b && b <= N")
			}
		} else {
			break
		}
		count += 1
	}

	for i := 0; i < N; i++ {
		fmt.Fprintf(writer, "%d ", baguniNum[i])
	}
}
