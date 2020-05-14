//use append function to append to a slice

//the first parameter of append is the slice followed
//by the values to be apended.

package main

import "fmt"

func main() {

	s := make([]int, 0, 5) // len = 0, cap = 5
	fmt.Println(s == nil)  //false

	s = append(s, 1, 2, 3, 4, 5, 6) // array is reallocated
	fmt.Println(s)                  //append works in zeroed slice

	var t []int
	fmt.Println(t == nil) //true
	t = append(t, 1, 2, 3, 4, 5, 6)
	fmt.Println(t) //append also works in nill slices

}
