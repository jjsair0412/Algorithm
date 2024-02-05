package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if i := 10; i < 11 {
		fmt.Println("test success")
	}

	// fmt.Println("i가 존재하나요 ? ", i)

	var a string = "test1"

	fmt.Printf("a : %v\n", a)

	a = "test2"

	fmt.Printf("re index a : %v\n", a)

	const (
		A = iota * 20
		B
		_
		D
	)

	fmt.Printf("A : %d\n", A)
	fmt.Printf("B : %d\n", B)
	fmt.Printf("D : %d\n", D)

	sum, i := 0, 0

	for i <= 100 {
		//  Error
		// sum += i++

		sum += i
		i++
	}

	//  Error
	// for j := 10; j >= 0; j--
	// {
	// 	fmt.Println("ex2 : ", j)
	// }

	for j := 3; j >= 0; j-- {
		fmt.Println("ex2 : ", j)
	}

	for ii := 0; ii < 5; ii++ {
		fmt.Println("ii : ", ii)
	}

	var (
		one int
		two int
	)

	reader := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscanf(reader, "%d %d\n", &one, &two)
	if err != nil {
		fmt.Println(os.Stderr, "Fscanf: %v\n", err)
	}

	if one < two {
		fmt.Println("<")
	} else if one > two {
		fmt.Println(">")
	} else if one == two {
		fmt.Println("==")
	}

}
