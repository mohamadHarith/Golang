//New returns a Value representing a pointer to a new zero value for the specified type.
//That is, the returned Value's Type is PtrTo(typ).

package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int `tag1:"First Tag" tag2:"Second Tag"`
	B string
}

func main() {
	f := Foo{A: 10, B: "Salutations"}

	fVal := reflect.New(reflect.TypeOf(f)) //creates a pointer to zero value of Foo

	fVal.Elem().Field(0).SetInt(5) //set values
	fVal.Elem().Field(1).SetString("hello")

	fmt.Println("value of fVal ", fVal.Elem())
	fmt.Println("type of fVal before calling Interface() ", reflect.TypeOf(fVal).Name())                   // returns Value
	fmt.Println("type of fVal after calling Interface() ", reflect.TypeOf(fVal.Elem().Interface()).Name()) // returns Foo
}
