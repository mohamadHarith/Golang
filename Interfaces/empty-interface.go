// The interface type that specifies zero methods is known as the empty interface:

// interface{}
// An empty interface may hold values of any type. (Every type implements at least zero methods.)

// Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.

package main

import "fmt"

type HelloWorldInterface interface {
	print()
}

type HelloWorldConcrete struct {
	s string
}

func (h HelloWorldConcrete) print() {
	fmt.Println(h.s)
}

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	var test HelloWorldConcrete
	test = HelloWorldConcrete{"My name is harith"}
	i = test
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
