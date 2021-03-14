package main

import (
	"fmt"
)


func main() {
	// atribuição simples
	type hotdog string

	x := 0
	var y hotdog = "James Bond"
	var z bool = false

	// print`s 
	fmt.Printf("a) %d\n", x)
	fmt.Printf("b) %s\n", y)
	fmt.Printf("c) %t\n", z)
	
	//verificação de tipo
	fmt.Printf("\nType var x equals %T\n", x)	
}