package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		N int
		X int
	)
	reader := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscanf(reader, "%d %d\n", &N, &X)
	if err != nil {
		fmt.Println("N, X 입력 에러")
		fmt.Println(err)
		return
	}

	slice := make([]int, N)

	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSuffix(numStr, "\n")
	numStr = strings.TrimSuffix(numStr, "\r")
	nums := strings.Split(numStr, " ")
	if len(nums) != N {
		fmt.Println("입력된 정수값 개수가 맞지 않습니다.")
		return
	}

	for cnt, Anum := range nums {
		slice[cnt], err = strconv.Atoi(Anum)
		if err != nil {
			fmt.Println("String convertor 에러")
			fmt.Println(err)
			return
		}
	}

	for i := range slice {
		if slice[i] < X {
			fmt.Printf("%d ", slice[i])
		}
	}
}
