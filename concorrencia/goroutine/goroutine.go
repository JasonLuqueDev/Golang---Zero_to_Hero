package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (Iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {

	// fale("Maria", "Pq voce não fala comigo ? ", 3)
	// fale("João", "Só posso falar depois de vc! ", 1)

	// go fale("Maria", "Ei... Sou Maria", 500)
	// go fale("João", "Opa... Sou João", 500)

	go fale("Maria", "Sou Maria", 10)
	fale("João", "Sou Joãoooo", 5)

}
