package main

import "fmt"

// a Função init é chamada antes de qualquer outra função em GO
// convenção Go

func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")

}
