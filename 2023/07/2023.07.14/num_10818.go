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
		N   int
		min int
		max int
	)

	reader := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscanf(reader, "%d\n", &N)
	if err != nil {
		fmt.Println("N 입력 에러")
		fmt.Println(err)
		return
	}

	slice := make([]int, N)

	numStr, _ := reader.ReadString('\r')
	numStr = strings.TrimSuffix(numStr, "\n")
	numStr = strings.TrimSuffix(numStr, "\r")

	arrStr := strings.Split(numStr, " ")
	if len(arrStr) != N {
		fmt.Println("배열 길이가 N과 다릅니다.")
		return
	}

	for i, arrNumStr := range arrStr {
		slice[i], err = strconv.Atoi(arrNumStr)
		if err != nil {
			fmt.Println("문자 배열 숫자변환 에러")
			return
		}
	}

	for i, realNum := range slice {
		if i == 0 {
			min = slice[0]
		} else {
			if min > realNum {
				min = realNum
			}
		}
	}
	for i, realNum := range slice {
		if i == 0 {
			max = slice[0]
		} else {
			if max < realNum {
				max = realNum
			}
		}
	}
	fmt.Printf("%d %d", min, max)

}
