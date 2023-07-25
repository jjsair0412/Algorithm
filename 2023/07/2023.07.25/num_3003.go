package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	inputByte, _, err := reader.ReadLine()
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("inputByte 에러발생")
		fmt.Println(err)
		return
	}

	arrStr := strings.Split(strings.TrimSpace(string(inputByte)), " ")

	arrInt := make([]int, len(arrStr))

	for i, strNum := range arrStr {
		arrInt[i], err = strconv.Atoi(strNum)
		if err != nil {
			fmt.Println("Atoi 에러발생")
			fmt.Println(err)
			return
		}
	}

	comChes := []int{1, 1, 2, 2, 2, 8}

	result := make([]int, len(comChes))

	for i, num := range comChes {
		result[i] = num - arrInt[i]
	}

	for _, resultNum := range result {
		fmt.Printf("%d ", resultNum)
	}
}
