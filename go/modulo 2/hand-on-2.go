package main

import "fmt"

func main()  {
	value := 5
	//adicionado 1 bit no final do valor
	aux := value << 1

	fmt.Printf("value: %d \n", value)
	fmt.Printf("Binary: %b \n",value)
	fmt.Printf("Hexa: %#x \n", value)
	fmt.Printf("Type: %T \n\n",value)
	
	fmt.Printf("value: %d \n", aux)
	fmt.Printf("Binary: %b \n",aux)
	fmt.Printf("Hexa: %#x \n", aux)
	fmt.Printf("Type: %T \n",aux)
	
}