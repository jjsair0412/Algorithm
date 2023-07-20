package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	input, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("input 입력에러 발생")
		fmt.Println(err)
		return
	}

	var alphabets = [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	for _, alphaByte := range alphabets {
		fmt.Printf("%d ", strings.IndexByte(string(input), alphaByte))
	}
}
