package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		culNum int // 출석번호
	)

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	allStu := make([]int, 30)

	for i := 0; i < 30; i++ {
		allStu[i] = i + 1 // 전체학생 번호 쓰기
	}

	for i := 0; i < 28; i++ {
		_, err := fmt.Fscanf(reader, "%d\n", &culNum)
		if err != nil {
			fmt.Println("culNum 입력에러 발생")
			fmt.Println(err)
			return
		}
		if contains(culNum, allStu) == true {
			allStu = remove(culNum, allStu)
			if allStu == nil {
				fmt.Println("변환 에러")
			}
		}
	}

	for i, _ := range allStu {
		if i+1 == len(allStu) {
			break
		} else {
			if allStu[i] > allStu[i+1] {
				tmp := 0
				tmp = allStu[i]
				allStu[i] = allStu[i+1]
				allStu[i+1] = tmp
			}
		}
	}

	for j, _ := range allStu {
		fmt.Printf("%d\n", allStu[j])
	}
}

func contains(s int, a []int) bool {
	for _, realNum := range a {
		if realNum == s {
			return true
		}
	}
	return false
}

func remove(i int, s []int) []int {
	for j, num := range s {
		if num == i {
			s[j] = s[len(s)-1]
			return s[:len(s)-1]
		}
	}
	return nil
}
