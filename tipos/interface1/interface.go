package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas IMPLICITAMENTE

func (p pessoa) toString() string {

	return p.nome + " " + p.sobrenome

}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}
func main() {
	var qualquerCoisa imprimivel = pessoa{"Jason", "Luque"}
	fmt.Println(qualquerCoisa.toString())
	imprimir(qualquerCoisa)

	qualquerCoisa = produto{"Camisa T", 70.00}
	fmt.Println(qualquerCoisa.toString())
	imprimir(qualquerCoisa)

	p2 := produto{"Calça Jeans", 189.99}
	imprimir(p2)

}
