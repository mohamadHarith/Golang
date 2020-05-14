//An array has a fixed size. A slicce, on the other hand is a dynamically-sized
//flexible view into the elements of an array. In practice, slices are much more
//common than arrays.

//The type []T is a slice with elements of type T.package Arrays

//A slice is formed by specifying two indices, a low and high bound,
//separated by a colon.

//a[low:high]

//This selects a half-open range which includes the first element, but excludes
//the last one.

package main

import "fmt"

func main() {
	//var a [6]int{2,3,5,7,11,13}
	var a [6]int
	a[0] = 2
	a[1] = 3
	a[2] = 5
	a[3] = 7
	a[4] = 11
	a[5] = 13
	fmt.Println(a)

	s := a[0:3]
	var t []int = a[0:4]
	fmt.Println(s)
	fmt.Println(t)
}
