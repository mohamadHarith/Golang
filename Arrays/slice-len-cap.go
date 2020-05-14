//a slice has both length and capacity

//the length of a slice is the number of elements it contains

// the capacity of a slice is the size of the underlying array

//a slice's length can be extended by resizing it, provided that it has sufficient capacity.

package main

import "fmt"

func main() {
	s := []int{1, 3, 5, 7, 11, 13}
	fmt.Println(s)

	//extend its length
	s = s[:4]
	fmt.Println(s)
}
