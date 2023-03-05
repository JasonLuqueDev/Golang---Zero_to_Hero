package main

import "fmt"

func main() {

	// array tem tipos homogêneos (mesmo tipo) e estática (fixo)
	// array é indexado, acessado através de um índice a partir de '0'

	var notas [3]float64 // notas contém 3 posições do tipo float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	// notas[3] = 10
	fmt.Println(notas)

	// Pegar a média das notas usando o FOR

	total := 0.0

	//
	for i := 0; i < len(notas); i++ {

		//
		total += notas[i]

	}

	// 'total' é do tipo float64, aqui converto o array em float64
	media := total / float64(len(notas))

	fmt.Printf("Média %.2f\n", media)

}
