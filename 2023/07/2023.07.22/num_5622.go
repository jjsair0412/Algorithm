package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var result int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	byteint, _, err := reader.ReadLine()
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("byteint 입력에러")
		fmt.Println(err)
		return
	}

	var alphabets = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	for _, num := range byteint {

		for j, alpha := range alphabets {
			if string(num) == alpha {
				result += cal(j)
			}
		}
	}

	fmt.Println(result)
}

func cal(j int) int {
	if j >= 0 && j <= 2 {
		return 3
	} else if j >= 3 && j <= 5 {
		return 4
	} else if j >= 6 && j <= 8 {
		return 5
	} else if j >= 9 && j <= 11 {
		return 6
	} else if j >= 12 && j <= 14 {
		return 7
	} else if j >= 15 && j <= 18 {
		return 8
	} else if j >= 19 && j <= 21 {
		return 9
	} else if j >= 22 && j <= 25 {
		return 10
	} else {
		return 0
	}
}
