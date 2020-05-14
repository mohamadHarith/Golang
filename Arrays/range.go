//the range form of the for loop iterates
//over a slice or map

//ranging ove slice return index and value

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Println(i, v)
	}

	//under score can be used to skip index or value
	// 	for i, _ := range pow
	// for _, value := range pow

}
