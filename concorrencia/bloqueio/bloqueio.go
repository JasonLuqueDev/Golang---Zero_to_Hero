package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // Operação bloqueante
	fmt.Println("Só executa depois que foi lido")
}

func main() {

	c := make(chan int) // Canal sem Buffer

	go rotina(c)
	fmt.Println(<-c) // Operação Bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c) //DeadLock
	fmt.Println("Fim")

}
