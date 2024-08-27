package main

import "fmt"

func main() {
	var a = map[int]string{}
	a[1] = "vishu"
	a[3] = "krishna"

	_, value := a[5]
	fmt.Println(value)
	for k, v := range a {
		fmt.Println(k, v)
	}
}
