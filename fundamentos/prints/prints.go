package main

import (
	"fmt"
)

func main() {
	fmt.Print("Mesma ")
	fmt.Print("Linha.")

	fmt.Println(" Nova")

	fmt.Println("linha.")

	x := 3.141516

	// Criado uma variavel 'xs' para converter a
	// variavel 'x' usando o Sprint

	xs := fmt.Sprint(x)

	// Imprime a variavel 'xs' convertida em numérico
	fmt.Println("O valor de x é " + xs)

	// Imprime numeros
	fmt.Println("O valor de x é", x)

	//  Imprime o valor usando 'DUAS' casas decimais
	// %.2f = Imprime variavel tipo FLOAT
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.99999
	c := false
	d := "opa"
	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)

}
