package main

import "fmt"

func printNum(num int) {
	if num == 0 {
		return
	}
	printNum(num - 1)
	fmt.Println(num)
}

//finding factorial..

func fact(num int) int {
	if num <= 1 {
		return 1
	}
	return num * fact(num-1)
}

func main() {
	var (
		a int = fact(4)
	)
	println(a)
}
