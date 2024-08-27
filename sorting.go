package main

import "fmt"

func bubbleSort(arr []int) {
	var x = len(arr)

	for i := 0; i < x-1; i++ {
		var flag bool = false
		for j := 0; j < x-i-1; j++ {
			if arr[j] > arr[j+1] {
				var temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
}

func main() {
	arr := []int{1, 2, 73, 42, 45, 1, 4}
	bubbleSort(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
