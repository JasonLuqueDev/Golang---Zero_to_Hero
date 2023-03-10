package main

import (
	"fmt"
	"reflect"
)

var medias [5]int

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {

	fmt.Printf("Média: %.2f", media(7.7, 3.2, 7.8, 2.3))
	fmt.Printf("\n ")
	fmt.Println(reflect.TypeOf(medias))
}
