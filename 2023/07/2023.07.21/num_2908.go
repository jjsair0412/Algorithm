package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	readStr, _, err := reader.ReadLine()
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("readStr 에러 발생")
		fmt.Println(err)
	}

	num1, num2 := makeIntArr(readStr)

	num1Result := makeNum(reverse(num1))

	num2Result := makeNum(reverse(num2))

	if num1Result > num2Result {
		fmt.Println(num1Result)
	} else {
		fmt.Println(num2Result)
	}
}

func makeIntArr(a []byte) ([]int, []int) {
	num1 := make([]int, 3)
	num2 := make([]int, 3)

	for i, str := range a {
		if i >= 0 && i <= 2 {
			num1[i], _ = strconv.Atoi(string(str))
		} else if i >= 4 && i <= 6 {
			num2[i-4], _ = strconv.Atoi(string(str))
		}
	}
	return num1, num2
}

func reverse(a []int) []int {
	var tmp int
	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j > i; j-- {
			tmp = a[j-1]
			a[j-1] = a[j]
			a[j] = tmp
		}
	}
	return a
}

func makeNum(a []int) int {
	var result int
	for i, num := range a {
		if i == 0 {
			result += num * 100
		} else if i == 1 {
			result += num * 10
		} else {
			result += num * 1
		}
	}
	return result
}
