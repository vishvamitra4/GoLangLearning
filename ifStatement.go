package main

import "fmt"

func main() {
	x := 18

	if x > 20 {
		fmt.Println("X is greater than 20..")
	} else if x <= 20 && x >= 10 {
		fmt.Println("X is less than 20 and greater than 10.")
	} else {
		fmt.Println("X is less than 10")
	}

}
