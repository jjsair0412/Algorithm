package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 가장 처음 바구니에는 공이 들어있지 않으며, 바구니에는 공을 1개만 넣을 수 있다.
func main() {
	var (
		N int // 공 , 바구니 번호
		M int // 공을 바구니에 넣을 횟수
	)

	reader := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscanf(reader, "%d %d", &N, &M)
	if err != nil {
		fmt.Println("N , M 입력에러 발생")
		fmt.Println(err)
		return
	}

}
