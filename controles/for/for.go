package main

import "fmt"

func main() {

	// Laço FOR por EXPRESSÃO

	i := 1
	for i <= 10 { //  Não tem While em GO
		fmt.Println(i)
		i++

	}

	// Laço FOR tradicional

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	// Achando valores pares/impares usando o módulo

	fmt.Println("\n Misturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}
}
