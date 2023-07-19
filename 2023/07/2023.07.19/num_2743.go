package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		str string
	)

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	_, err := fmt.Fscanf(reader, "%s\n", &str)
	if err != nil {
		fmt.Println("str 입력에러 발생")
		fmt.Println(err)
		return
	}

	fmt.Fprintf(writer, "%d\n", len(str))
}
