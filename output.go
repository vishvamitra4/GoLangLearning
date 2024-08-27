package main

import "fmt"

/*

func main() {
	var name = "vishvamitra"
	var age = "two three"

	fmt.Println(name, "\n", age)
}
*/

// Printf in golang...

func main() {
	var i string = "hello world!"
	var j float32 = 34.454

	fmt.Printf("values of i is %v and type is %T\n", i, i)
	fmt.Printf("values of j is %v ,type is %T\n", j, j)
	fmt.Printf("%.2f", j)
}
