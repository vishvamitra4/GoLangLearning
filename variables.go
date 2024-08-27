package main

import "fmt"

/*
// all variable get initialed at with its declaration..
func main() {
	var a string
	var b float32
	var c bool
	var d int

	fmt.Println(a, b, c, d)
}



func main() {
	var name string = "vishvamitra"
	// type is inferred..
	// short declaration oprator..
	myName := "krishna singh"
	fmt.Println(name)
	fmt.Println(myName)
}



// go multiple variable declaration..

func main() {
	var a, b, c int = 1, 4, 5
	fmt.Println(a, b, c)
}


func main() {
	var a, b = 6, "Hello wolrd"
	c, d := 7, "World!"

	fmt.Println(a, b, c, d)

}


// go varialbe declaration in block..

func main() {
	var (
		a int
		b int    = 1
		c string = "Hello"
	)

	fmt.Println(a, b, c)
}
*/

// declare the const values..

func main() {
	const (
		A int    = 1
		B string = "Vishvamitra"
	)

	fmt.Println(A, B)
}
