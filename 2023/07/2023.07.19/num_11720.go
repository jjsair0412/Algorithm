package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	// "strings"
)

func main() {
	var (
		N        int
		inputstr string
	)
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanf(reader, "%d\n", &N)

	slice := make([]string, N)

	fmt.Fscanf(reader, "%s", &inputstr)

	fmt.Printf("%v\n", slice)

}
