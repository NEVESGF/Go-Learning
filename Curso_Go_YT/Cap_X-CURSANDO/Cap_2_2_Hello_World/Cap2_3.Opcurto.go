package main

import (
	"fmt"
)

// para definir uma função a nível de package deve-se utilizar o var não a marmota
// e.g. var y = "Olá, bom dia!"

func main() {

	// := deve ser utilizado para declarar pelo menos uma variável
	// A marmota (:=) só funciona dentro de um code block
	// = deve ser utilizado para definir um novo valor para uma variável já declarada

	x := 10
	y := "Olá, bom dia"

	//Existem KeyWords (palavras reservadas) essas palavras não podem ser utilziadas para definir constantes
	// e.g func, packge, import

	fmt.Printf("x: %v, %T\n", x, x) // %v (value) e %T (type) quando após o , x x é como se fosse um format do python
	fmt.Printf("y: %v, %T", y, y)

	
}
