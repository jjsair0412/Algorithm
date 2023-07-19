package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		T   int
		str string
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanf(reader, "%d\n", &T)

	for i := 0; i < T; i++ {
		fmt.Fscanf(reader, "%s\n", &str)

		fmt.Printf("%s%s\n", string(str[0]), string(str[len(str)-1]))
	}
}
