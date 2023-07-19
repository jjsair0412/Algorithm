package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		str byte
	)

	reader := bufio.NewReader(os.Stdin)

	str, _ = reader.ReadByte()
	fmt.Println(str)
}
