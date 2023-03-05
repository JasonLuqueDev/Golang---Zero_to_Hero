package main

import "fmt"

// func nomeDaFuncao
// não recebe dados e não retorna nada
func f1() {
	fmt.Println("F1")

}

// função que recebe dois parametros tipo string e não retorna nada
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

// função que não recebe nada e 'RETORNA' uma string
func f3() string {
	// o RETURN é obrigatório
	return "F3"
}

// é possivel declarar assim funções do mesmo tipo, ex. string
// neste exemplo a função retorna tbm string
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)

}

// função que permite receber multiplos valores
func f5() (string, string) {
	return "Retorno 1", "Retorno2"
}

func main() {
	f1()
	f2("Param1", "Param2")

	r3, r4 := f3(), f4("PARAM1", "PARAM2")
	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := f5()
	fmt.Println("F5: ", r51, r52)

}
