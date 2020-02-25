//The var statement declares a list of variables; as in function argument lists, the type is last.

package main

import "fmt"

//global variable
var c, python, java bool //multiple variable declaration of type bool

func main() {
	//local variable
	var i int //if unintialized the value is 0
	fmt.Println(i, c, python, java)
}
