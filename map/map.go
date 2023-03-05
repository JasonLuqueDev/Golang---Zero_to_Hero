package main

import "fmt"

func main() {
	// var aprovados map[int]string

	// Map DEVEM ser inicializados

	aprovados := make(map[int]string)
	aprovados[12312312300] = "Maria"
	aprovados[23423423400] = "Pedro"
	aprovados[99988877766] = "Ana"

	fmt.Println(aprovados)
	//

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	// Imprime um valor da chave
	fmt.Println(aprovados[12312312300])

	// Deletar um valor do Map
	delete(aprovados, 12312312300)
	fmt.Println(aprovados)

}
