//A function with pointer argument must take a pointer
//Ex:
//var v Vertex
//ScaleFunc(v, 5) //compile error
//ScaleFunc(&v, 5)// OK

//A method with pointer receiver may take both value or pointer as its receiver
//Ex:
//v.Scale(5) //OK
//p := &v
//p.Scale(5) //OK

// For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.

package main

import "fmt"

type Vertex struct {
	x, y float64
}

func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func (v Vertex) Scale2(f float64) Vertex {
	v.x = v.x * f
	v.y = v.y * f
	return v
}

func ScaleFunc(v *Vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := Vertex{3, 4}

	v.Scale(2)
	ScaleFunc(&v, 10)
	fmt.Println(v)

	p := &Vertex{3, 4}
	p.Scale(3)
	ScaleFunc(p, 10)
	fmt.Println(*p)

	//equivalent thing happens in reverse.
	//a function that takes value argument must take value
	//a method that takes value receiver take either a value or a pointer
	a := Vertex{3, 2}
	fmt.Println(a.Scale2(10))
	b := &a
	fmt.Println((*b).Scale2(2))
	fmt.Println(b.Scale2(2)) //intepreted as the above
}
