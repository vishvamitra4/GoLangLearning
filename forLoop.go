package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	var sum = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i] * 2
	}

	fmt.Println(sum)
}
