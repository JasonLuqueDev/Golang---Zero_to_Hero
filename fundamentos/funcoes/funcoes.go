package main

import "fmt"

// função somar, recebe 'a' e 'b' como inteiro
// e também retorna um inteiro 'int'
func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {
	resultado := somar(3, 4)
	imprimir(resultado)

}
