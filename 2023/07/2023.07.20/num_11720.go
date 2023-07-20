package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		N         int
		inputstr  string
		addResult int
	)
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanf(reader, "%d\n", &N)

	fmt.Fscanf(reader, "%s", &inputstr)

	slice := strings.Split(inputstr, "")

	for _, num := range slice {
		val, _ := strconv.Atoi(num)
		addResult += val
	}

	fmt.Println(addResult)
}
