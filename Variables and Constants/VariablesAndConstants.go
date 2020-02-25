package main 

import "fmt"

func main(){
	//Constant
	const pi float64 = 3.14159

	//variable
	var A int = 3

	//multiple variable declaration
	var (
		B = 4
		C = 5
	)

	var D string = "Harith"

	fmt.Println(A)
	fmt.Println(B, C)
	fmt.Println(D)
	//string length
	fmt.Println(len(D))

}