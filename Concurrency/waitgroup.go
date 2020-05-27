package main

import (
	"fmt"
	"time"
)

func main() {

	//this will exit because the main goroutine (main function exits when there is no lines to
	//be executed after the below go routines.
	//)
	go count("sheep")
	go count("fish")

	//can be locked using sleep or scanline (blocking call)
	time.Sleep(time.Millisecond * 500)
	fmt.Scanln()

	//can be better handled using wait group

}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
