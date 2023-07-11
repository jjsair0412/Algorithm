package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		count int
		A     int
		B     int
	)

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(reader, &count)

	for i := 0; i < count; i++ {
		fmt.Fscanln(reader, &A, &B)
		fmt.Fprintln(writer, A+B)
	}

	writer.Flush()
}
