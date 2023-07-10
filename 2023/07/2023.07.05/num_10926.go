package main

import "fmt"

func main() {
	var (
		id        string
		id_backup string
	)

	fmt.Scanf("%s", &id)

	id_backup = id

	if id == id_backup {
		fmt.Println(id + "??!")
	}
}
