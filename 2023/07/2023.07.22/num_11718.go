package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		inputByte, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("inputByte 입력에러")
			fmt.Println(err)
			return
		}
		fmt.Println(string(inputByte))
	}

}
