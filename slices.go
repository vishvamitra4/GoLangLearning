package main

// import a particular package..
import "fmt"

/*
slice_name := []datatype{values}
slice_name := make([]type , length , capacity)
*/

// GO SLICES

func main() {
	your_slice := make([]int, 5, 51)
	fmt.Println(len(your_slice))
	fmt.Println(cap(your_slice))
	fmt.Println((your_slice))
}
