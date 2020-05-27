package main

import (
	"fmt"
	"time"
)

func main() {

	//	the first call will loop infinitely with calling the second call
	// count("sheep")
	// count("fish")

	// go routine - run the call preceeding the go in the background and execute the next line immediately
	go count("sheep")
	count("fish")
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
