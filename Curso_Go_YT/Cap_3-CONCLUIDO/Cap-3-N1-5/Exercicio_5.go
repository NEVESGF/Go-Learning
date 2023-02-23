package main

import (
	"fmt"
)

type exercisefour int

var x exercisefour

var y int

func main() {

	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)
	y = int(x)
	fmt.Print(y)
	fmt.Printf("\n%T", y)

}
