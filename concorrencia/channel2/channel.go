package main

import (
	"fmt"
	"time"
)

// Channel (canal) - É a forma de comunicação entre GoRoutines
// channel = tipo

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // Enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base

}

func main() {

	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c // Recebendo os dados do canal
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)
}
