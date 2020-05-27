//channel is like a pipe that can be used to send and receive message
//sending and receiving messages through a channel are blocking operations

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string) //	make a channel
	go count("sheep", c)   //	pass the channel

	//this will only receive one message and terminate everything
	// msg := <-c //	blocks and waits to receive a message
	// fmt.Println(msg)

	// to receive multiple messages
	// for {
	// 	msg, open := <-c
	// 	fmt.Println(msg) // this will create a deadlock if the go routine doesnt close the channel after finish executing
	// 	if !open {
	// 		break
	// 	}
	// }

	//short form for the above (syntactic sugar)
	for msg := range c {
		fmt.Println(msg)
	}

}

func count(thing string, c chan string) { //channel can be of any type
	for i := 1; i <= 5; i++ {
		c <- thing //blocks and waits for receiver to be ready and sends a message
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
