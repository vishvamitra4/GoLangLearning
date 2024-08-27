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
}

// here i have added main branch..
