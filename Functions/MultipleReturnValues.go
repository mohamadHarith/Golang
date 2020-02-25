package main

import "fmt"

func swap(x, y string) (string, string) { //multiple return types
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}