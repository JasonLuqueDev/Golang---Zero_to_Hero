package main

import "fmt"

type curso struct {
	nome string
}

func main() {

	var interfaceVazia interface{}

	fmt.Println(interfaceVazia)

	interfaceVazia = 3
	fmt.Println(interfaceVazia)

	type dinamico interface {
	}

	var interfaceVazia2 dinamico = "Opa"
	fmt.Println(interfaceVazia2)

	interfaceVazia2 = curso{"Golang"}

}
