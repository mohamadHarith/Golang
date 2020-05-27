package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup //counter
	wg.Add(1)

	//self invoking anonymous function
	go func() {
		count("sheep")
		wg.Done() //decrements the counter
	}()

	wg.Wait() // this will block the main go routine from exiting until the counter is zero

}

func count(thing string) {
	for i := 1; i < 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
