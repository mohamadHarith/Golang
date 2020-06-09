package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "greetings"

	gval := reflect.ValueOf(s) // returns a variable of type reflect.Value

	fmt.Println(reflect.TypeOf(gval).Name()) //returns Value

	fmt.Println(reflect.TypeOf(gval.Interface()).Name()) //Interface returns v's current value as an interface{}.
	//It is equivalent to:
	//var i interface{} = (v's underlying value)

	// gval.SetString("hello") //will panic becase cant set variable passed as value
	// fmt.Println(s)

	gvalp := reflect.ValueOf(&s)
	gvalp.Elem().SetString("hello") //Elem() dereferences the pointer
	fmt.Println(s)

}
