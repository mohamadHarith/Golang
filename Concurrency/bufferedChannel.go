package main 

import (
	"fmt"
)

func main(){
	// c := make(chan string)

	// c <- "hello" // this will produce deadlock bcs it cant continue to the next line

	// msg := <- c
	// fmt.Println(msg)

	c := make(chan string, 2) //creates a channel buffer of size 2 so that the lines can execute until 2 messages are received

	c <- "hello"
	c <- "world"
	//c <- "test" //chanel is full

	msg := <- c
	fmt.Println(msg)

	msg = <- c
	fmt.Println(msg)
}