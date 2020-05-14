//a struct is a collection of fields

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	//array of structs
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{1, false},
	}
	fmt.Println(s)
}
