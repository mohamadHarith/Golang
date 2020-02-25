//A pointer holds the memory address of a value.
//The & operator generates a pointer to its operand
//The * operator denotes the pointer's underlying value

package main

import (
	"fmt"
)

func main(){
	
	a := 6
	var p *int //pointer variable declaration
	p = &a //pointer variable p contains the memory address of variable a
	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p) // dereferencing - read a through pointer p
	*p = 21 // indirecting - set a through pointer p
	fmt.Println(a) // 21
	fmt.Println(*p)//21

}