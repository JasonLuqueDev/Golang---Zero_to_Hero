package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso

	// Pode adicionar mais métodos

}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo Ligado...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Fazendo Baliza Automaticamente Ligado...")

}

func main() {
	var b esportivoLuxuoso = bmw7{}
	b.ligarTurbo()
	b.fazerBaliza()

}
