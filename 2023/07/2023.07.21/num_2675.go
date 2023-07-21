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
		var count int
		for {
			if count == R-1 {
				fmt.Println()
				break
			} else {
				for h := 0; h < len(strArr); h++ {
					for j := 0; j < R; j++ {
						fmt.Printf("%s", strArr[h])
					}
				}
			}
			count += 1
		}
	}
}
