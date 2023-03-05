package main

import (
	"fmt"
	"reflect"
)

func main() {

	a1 := [3]int{1, 2, 3} // Array

	s1 := []int{1, 2, 3} // Slice = Tamanho variável []

	fmt.Println(a1, s1)
	//
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))
	//
	a2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a2)

	// Slice não é um Array!
	// Slice define um pedaço de um Array

	s2 := a2[1:3] // 2 3
	fmt.Println(a2, s2)
	s3 := a2[:2]
	fmt.Println(s3)

}
