package main

import "fmt"

func main() {
	m := make(map[string]int)

	//insert
	m["answer"] = 42
	fmt.Println(m)

	//update
	m["answer"] = 48
	fmt.Println(m)

	//delete
	delete(m, "answer")
	fmt.Println(m)

	v, ok := m["answer"]
	fmt.Println(v, ok) //if present ok true, if not present ok false v is the zero value for the maps element type
}
