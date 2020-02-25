// Type inference
// When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

// When the right hand side of the declaration is typed, the new variable is of that same type:

// var i int
// j := i // j is an int

package main

import "fmt"

func main() {
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}
