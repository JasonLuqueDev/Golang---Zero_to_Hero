package main

import (
	"fmt"
	"reflect"
)

// tipo NOME DO TIPO/IDENTIFICADOR struct
type produto struct {

	// agrupar dados,
	// PARECIDO com classe
	// identificador e o tipo
	nome     string
	preco    float64
	desconto float64
}

// Método: Função com 'receiver' (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	// definir produto
	var produto1 produto

	// inicializamos o produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Notebook", 1989.90, 0.10}
	fmt.Println(produto2.precoComDesconto())
	fmt.Println(reflect.TypeOf(produto2))
}
