package main

import (
	"bufio"
	"fmt"
	"os"
)

// 가장 처음 바구니에는 공이 들어있지 않으며, 바구니에는 공을 1개만 넣을 수 있다.
func main() {
	var (
		N     int // 공 , 바구니 번호
		M     int // 공을 바구니에 넣을 횟수
		count int
		a     int
		b     int
		c     int
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

	baguniNum := make([]int, N)

	for {
		if M != count {
			_, err := fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
			if err != nil {
				fmt.Println("a,b,c 입력에러발생")
				fmt.Println(err)
				return
			}
			for M := a; M <= b; M++ {
				baguniNum[M-1] = c
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
