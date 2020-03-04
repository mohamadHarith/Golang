//Go's return values may be named. If so, they are treated as variables defined at the top of the function.
//Naked return should only be used in short functions

package main

import "fmt"

func split(sum int) (x, y int) { //named return value must be enclosed in parentheses
	x = sum * 4 / 9
	y = sum - x
	return // this is known as naked return - return statement without arguments
}

func main() {
	fmt.Println(split(17))
}

