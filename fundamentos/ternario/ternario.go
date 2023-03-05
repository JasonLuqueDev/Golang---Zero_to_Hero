package main

import "fmt"

// NÃO EXISTE OPERADOR TERNÁRIO EM GO

func obterResultado(nota float64) string {

	// return nota >= 6 ? "Aprovado" : "Reprovado"

	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

}

func main() {
	fmt.Println(obterResultado(0.2))
}
