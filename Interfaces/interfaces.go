//An interface type is defined as a set of method signatures.

//A type (class / struct) implements an interface by implementing ALL its methods.

//There is no explicit declaration of intent, no implements keyword

package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

//type Rectangle implements the interface Rectangle
func (r Rectangle) perimeter() float64 {
	return 2*r.height + 2*r.width
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func measure(g Geometry) { //g is an interface value. calling method through g will call the method of the underlying concrete type
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	a := Rectangle{width: 20, height: 10}
	b := Circle{radius: 5}
	measure(a) //if a implements the methods of interface Geometry through pointer receiver, a should be passed as address, &a
	measure(b)
}
