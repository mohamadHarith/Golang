package main

//Variables declared without an explicit initial value are given their zero value.

// The zero value is:

// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println(i, f, b, s)
}
