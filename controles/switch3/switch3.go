package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {

	switch i.(type) {
	case int:
		return "Inteiro"

	case float32, float64:
		return "Real / Float"

	case string:
		return "String"

	case func():
		return "Função"

	case bool:
		return "Boolean"

	default:
		return "Tipo Inválido..."
	}

}

func main() {
	fmt.Println(tipo(true))
	fmt.Println(tipo(2222.455))
	fmt.Println(tipo("fdasfsdfa"))
	fmt.Println(tipo(time.Now()))
}
