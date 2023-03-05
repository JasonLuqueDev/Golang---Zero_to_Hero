package main

import "fmt"

// uma variável que recebe uma função anônima
var soma = func(a, b int) int {
	return a + b
}

var sub = func(a, b int) int {
	return a - b
}

func main() {
	div := func(a, b int) int {
		return a / b
	}

	fmt.Println(soma(2, 3))
	fmt.Println(sub(2, 3))
	fmt.Println(div(2, 3))

}
