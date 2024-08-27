package main

import "fmt"

func getMsg() {
	fmt.Println("Line got executed..")
}

func printName(name string) string {
	var yourName = name
	return yourName
}

// printing the arrat and modify it...
func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i], " ")
	}
}

// printing sum of all element..
func printSum(arr []int) int {
	var sum int = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		arr[i] *= sum
	}

	return sum
}

func main() {
	arr := []int{2, 434, 5}
	fmt.Println(printSum(arr))
	printArray(arr)
}
