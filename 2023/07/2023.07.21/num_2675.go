package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		T int
		R int
		S string
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanf(reader, "%d\n", &T)

	for i := 0; i < T; i++ {
		fmt.Fscanf(reader, "%d %s\n", &R, &S)
		strArr := strings.Split(S, "")
		for _, str := range strArr {
			for j := 0; j < R; j++ {
				fmt.Printf("%s", str)
			}

		}
		fmt.Println()
	}
}
