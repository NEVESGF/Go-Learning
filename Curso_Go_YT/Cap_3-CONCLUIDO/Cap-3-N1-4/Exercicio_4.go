package main

import (
	"fmt"
)

type exercisefour int
var x exercisefour

func main() {

	fmt.Printf("%v\n",x)
	fmt.Printf("%T\n",x)

	x = 42

	fmt.Printf("%v\n",x)
}