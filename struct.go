// Struct is used to create single variable using different variable..

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	// var Vishu Person
	// Vishu.name = "Vishvamitra"
	// Vishu.age = 234
	// fmt.Println(Vishu.name, Vishu.age)

	var Vishu *Person = new(Person)
	Vishu.name = "Vishvamitra"
	fmt.Println(&Vishu, &Vishu.name, &Vishu.age)
}
