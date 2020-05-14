//slices can be created with the buitl in make function.package Arrays

//this is how dynamically-sized arrays are created.

//the make function allocates a zeroed array and return
// a slice that refers to that array

package main

import "fmt"

func main() {

	a := make([]int, 5) //len(a)=5
	fmt.Println(a)

	b := make([]int, 0, 5) //len(b)=0 cap(b)=5
	fmt.Println(b)

	c := b[:2]
	fmt.Println(c)

	d := c[2:5]
	fmt.Println(d)

}
