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

	fmt.Fscanf(reader, "%d\n", &count)

	for i := 0; i < count; i++ {
		fmt.Fscanf(reader, "%d %d\n", &A, &B)
		fmt.Fprintf(writer, "Case #%d: %d\n", i, A+B)
	}

	writer.Flush()
}
