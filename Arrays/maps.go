//a map maps keys to value.

//the zero value of a map is nil. A nil map
//has no keys nor can keys be added.package Arrays

//the make function returns a map of the given
//type, initalized and ready for use.

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//a variable m of type map whose key is a string and value is a Vertex
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}
