//A slice literal is like an array literal without the length

package main

import "fmt"

func main() {

	//this is an array literal
	p := [6]int{2, 3, 4, 7, 11, 13}

	//creates the same array as above, then builds a slice that references it
	q := []int{2, 3, 4, 7, 11, 13}

	fmt.Println(p, q)

}
