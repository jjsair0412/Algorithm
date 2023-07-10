package main

import "fmt"

func main() {
	var (
		first  int
		second int
	)
	fmt.Scanf("%d %d", &first, &second)

	fmt.Printf("%d\n", first+second)
	fmt.Printf("%d\n", first-second)
	fmt.Printf("%d\n", first*second)
	fmt.Printf("%d\n", first/second)
	fmt.Printf("%d\n", first%second)
}
