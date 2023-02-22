package main

import (
	"fmt"
)

var x int = 10  // tipos básicos: int, string, bool

func main() {

	//x = 20.5 - Mostrará um erro visto que quando feito o var foi definido x como int
	//quando declarado uma váriavel ela é imutavel, será como foi definida no começo até o fim

	fmt.Println(x)

}