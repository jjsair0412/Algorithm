package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 시간초과 발생 , 버블정렬 해결방법 찾아야함
func main() {
	var (
		N     int
		K     int
		count int
		swap  int
		A     int
		B     int
	)

	reader := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscanf(reader, "%d %d\n", &N, &K)
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

	for {
		if count != len(slice) {
			// 5
			// 0 1 2 3 4
			for i := range slice {
				if i+1 == len(slice) {
					break
				}
				if slice[i] > slice[i+1] {
					tmp := slice[i+1]
					slice[i+1] = slice[i]
					slice[i] = tmp
					swap += 1
					if swap == K {
						A = slice[i]
						B = slice[i+1]
					}
				}
			}
		} else {
			break
		}
		count += 1
	}
	if swap < K {
		fmt.Println("-1")
	} else {
		fmt.Printf("%d %d", A, B)
	}

}
