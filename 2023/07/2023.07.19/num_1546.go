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
	)

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	_, err := fmt.Fscanf(reader, "%d\n", &N)
	if err != nil {
		fmt.Println("N 입력에러")
		fmt.Println(err)
		return
	}

	slice := make([]int, N)

	numStr, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("숫자입력 에러발생")
		fmt.Println(err)
		return
	}

	arrStr := strings.Split(strings.TrimSpace(string(numStr)), " ")
	if len(arrStr) != len(slice) {
		fmt.Println("배열길이가 달라요.")
		return
	}

	for i, strNum := range arrStr {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			fmt.Println("Atoi 에러 발생")
			fmt.Println(err)
			return
		}
		slice[i] = num
	}

	M := maxFunc(slice)

	slice2 := make([]float32, N)
	for i, realNum := range slice {
		slice2[i] = float32(float32(realNum) / float32(M) * 100)
	}

	fmt.Fprintf(writer, "%f", average(slice2))
}

func maxFunc(slice []int) int {
	var max int
	for i, num := range slice {
		if i == 0 {
			max = num
		} else {
			if max < num {
				max = num
			}
		}
	}
	return max
}

func average(slice []float32) float32 {
	var AllAddNum float32
	for _, num := range slice {
		AllAddNum += num
	}

	return float32(AllAddNum / float32(len(slice)))
}
