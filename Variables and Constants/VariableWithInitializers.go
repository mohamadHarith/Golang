//A var declaration can include initializers, one per variable.
//If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

package main

import "fmt"

var i, j int = 1, 2 //follows order, one per variable

func main() {
	var c, python, java = true, false, "no!" //the type is omitted
	fmt.Println(i, j, c, python, java)
}