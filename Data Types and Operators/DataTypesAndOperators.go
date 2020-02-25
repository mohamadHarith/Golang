package main

import (
	"fmt"
)

func main(){
	
	//data types
	var a int = 5
	var b float32 = 4.32
	const pi float64 = 3.14151

	//shorthand notation for multiple variable declaration without type.
	//go handles the type
	x,y := 14,15

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(pi)
	fmt.Println(x, y)


	//operators
	i,j := 5,6
	fmt.Println("x+y = ", i+j)
	fmt.Println("x-y = ", i-j)
	fmt.Println("x*y = ", i*j)
	fmt.Println("x/y = ", i/j)
	fmt.Println("x mod y = ", i%j)

	//booleans
	//var isBool bool = true
	isBool := true
	hate := false

	fmt.Println(isBool && hate)
	fmt.Println(isBool || hate)
	fmt.Println(!isBool)

}