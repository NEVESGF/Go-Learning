package main

import(
	"fmt"
)

var a int
var b float64
var c string
var d bool

func main() {

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

	//Antes da variavel ser inicializada pelo usuário ela será valor zero
	// - ints: 0
	// - floats: 0.0
	// - booleans: false
	// - strings: ""
	// - pointers, functions, interfaces, slices, channels, maps: nil
}

