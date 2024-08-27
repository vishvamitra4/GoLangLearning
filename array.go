package main

import "fmt"

/*

func main() {
	var arr1 = [2]int{1, 4}
	// with short declaratio sign
	arr2 := [2]string{"vishu", "barua"}

	fmt.Println(arr1, arr2)
}
*/

// Initialization in array...

func main() {
	var a = [5]int{}               // not initialzed..
	var b = [5]int{1, 4, 5, 6}     // partial inti...
	var c = [5]int{1, 2, 3, 4, 5}  // fully init...
	var d = [5]int{1: 100, 4: 567} // init.. at particular index...

	fmt.Println(a, b, c, d)
	// length of array..
	fmt.Println(len(d))
}
