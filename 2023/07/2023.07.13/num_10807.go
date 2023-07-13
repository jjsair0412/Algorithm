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
		cnt     int
		findNum int
		sameCnt int
	)
	reader := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscanf(reader, "%d\n", &cnt)
	if err != nil {
		fmt.Println("cnt 입력 에러")
		fmt.Println(err)
		return
	}

	switch cnt >= 1 && cnt <= 100 {
	case true:

		slice := make([]int, cnt)

		numStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("정수값 입력 에러")
			fmt.Println(err)
			return
		}
		numStr = strings.TrimSuffix(numStr, "\n")
		numStr = strings.TrimSuffix(numStr, "\r")
		nums := strings.Split(numStr, " ")
		if len(nums) != cnt {
			fmt.Println("입력된 정수값 개수가 맞지 않습니다.")
			return
		}

		for i, num := range nums {
			val, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("정수값 변환 에러")
				fmt.Println(err)
				return
			}
			slice[i] = val
		}

		_, err = fmt.Fscanf(reader, "%d", &findNum)
		if err != nil {
			fmt.Println("findNum 입력 에러")
			fmt.Println(err)
			return
		}

		for i := range slice {
			if slice[i] == findNum {
				sameCnt += 1
			}
		}

		fmt.Println(sameCnt)
	case false:
		fmt.Println("1부터 100 사이의 값을 넣어주세요.")
	}

}
