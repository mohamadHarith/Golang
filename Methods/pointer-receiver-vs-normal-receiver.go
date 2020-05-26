package main

import (
	"fmt"
)

type point struct {
	x int
	y int
}

func (p *point) scale(w int) {
	p.x = p.x * w
	p.y = p.y * w
}

func (p point) scale2(w int) (q point) {
	//cant do this
	//p.x = p.x*w
	//p.y = p.y*w
	//return p

	q.x = p.x * w
	q.y = p.y * w

	return
}

func main() {
	fmt.Println("Hello, playground")

	p1 := point{2, 3}
	fmt.Println(p1)

	//method with pointer receiver
	p1.scale(10)
	fmt.Println(p1)

	//method with normal receiver
	fmt.Println(p1.scale2(3))

	fmt.Println(p1)
}
