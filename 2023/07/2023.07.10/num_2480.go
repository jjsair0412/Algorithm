package main

import "fmt"

func main() {
	var (
		one   int
		two   int
		thr   int
		same  int
		count int
		tmp   int
	)

	_, err := fmt.Scanf("%d %d %d", &one, &two, &thr)
	if err != nil {
		fmt.Println("입력 에러")
		return
	}

	var arr = [...]int{one, two, thr}

	for i := 0; i < len(arr); i++ {
		if i+1 == 3 {
			break
		}
		if arr[i] == arr[i+1] {
			count += 1
			same = arr[i]
		} else if i == 0 && arr[i] == arr[i+2] {
			count += 1
			same = arr[i]
		}
	}

	if count == 2 {
		fmt.Printf("%d", 10000+one*1000)
	} else if count == 1 {
		fmt.Printf("%d", 1000+same*100)
	} else {
		if arr[0] > arr[1] {
			tmp = arr[0]
			arr[0] = arr[1]
			arr[1] = tmp
		}
		if arr[1] > arr[2] {
			tmp = arr[1]
			arr[1] = arr[2]
			arr[2] = tmp
		}
		if arr[0] > arr[1] {
			tmp = arr[0]
			arr[0] = arr[1]
			arr[1] = tmp
		}

		fmt.Printf("%d", arr[2]*100)
	}
}
