package main

import "fmt"

func fun(s1 string, s2 string) (string, string) {
	var a = s1 + "fgg"
	var b = s2 + a
	return a, b
}

func main() {

	a, b := fun("Hey", "Worlf")
	fmt.Println(a, b)

	/*
		var arr = []int{1, 2, 3, 4, 5, 5, 3}
		var M = map[int]int{}

		for i := range arr {
			M[arr[i]]++
		}

		for k, v := range M {
			fmt.Println(k, v)
		}
	*/
}
