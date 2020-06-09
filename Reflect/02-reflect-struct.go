package main

import (
	"fmt"
	"reflect"
)

type structA struct {
	a int      `tag:"this is a tag`
	b []string `tag1:"this is a tag" tag2:"this is another tag"`
}

func main() {
	a := structA{a: 1, b: []string{"hello", "world"}}
	reflectVariable := reflect.TypeOf(a)

	fmt.Println(reflectVariable.Kind()) //returns the derived type struct

	fmt.Println(reflectVariable.NumField()) //returns the number of fields in the sturct

	//iterate over the struct fileds
	for i := 0; i < reflectVariable.NumField(); i++ {

		field := reflectVariable.Field(i) //returns ith filed

		fmt.Println(field.Type.Name()) //returns the field type. if it is derived type returns empty string

		fmt.Println(field.Type.Kind()) //return derived types

		if field.Tag != "" {
			fmt.Println(field.Tag)             //returns the tag in the string
			fmt.Println(field.Tag.Get("tag2")) //gets a specific tag
		}
	}
}
