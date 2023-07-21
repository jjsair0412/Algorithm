package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var count int

	reader := bufio.NewReader(os.Stdin)

	readStr, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("inbyte 입력 에러 발생")
		fmt.Println(err)
		return
	}
	text := strings.TrimSpace(readStr)

	resultArr := strings.Split(string(text), " ")

	for _, resultStr := range resultArr {

		if !strings.EqualFold(resultStr, "") && resultStr != "\n" {
			count += 1
		}
	}

	fmt.Println(count)
}
