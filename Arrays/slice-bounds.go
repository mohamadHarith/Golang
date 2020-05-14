//when slicing, you may omit the high or low bounds to use their
//defaults instead.package Arrays

//the default is zero for the low bound and the length of the slice for
//high bound

package main

import "fmt"

func main() {
	s := [6]int{2, 3, 5, 7, 11, 13}

	//all these are equivalent
	t := s[0:6]
	fmt.Println(t)

	t = s[:6]
	fmt.Println(t)

	t = s[:]
	fmt.Println(t)

}
