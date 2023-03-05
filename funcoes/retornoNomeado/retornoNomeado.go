package main

import "fmt"

// função que recebe 'dois inteiros' p1 e p2
// com retorno 'nomeado' primeiro e segundo
func trocar(p1 int, p2 int) (segundo int, primeiro int) {

	segundo = p2
	primeiro = p1
	return //retorno limpo
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)

}
