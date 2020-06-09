package main

import (
	"fmt"
	"reflect"
)

// Reflection in Go is built around three concepts: Types, Kinds, and Values.
//The reflect package in the standard library is the home for the types and functions that implement reflection in Go.

func test() {
	var a int = 1
	b := reflect.TypeOf(a) //returns a variable of type reflect.Type which contains many methods thath has info about the type

	fmt.Println(b.Name()) //returns the name of the type. If derived type, such as slice and struct, empty string is returned

	c := []string{"a", "b"}
	fmt.Println(reflect.TypeOf(c).Kind()) // return the derived type slice
}
