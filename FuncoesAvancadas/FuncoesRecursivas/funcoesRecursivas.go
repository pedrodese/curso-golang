package main

import (
	"fmt"
	"time"
)

// exemplo de função recursiva em GO
func retornaPosicaoNumeroFibonnaci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return retornaPosicaoNumeroFibonnaci(posicao-2) + retornaPosicaoNumeroFibonnaci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")
	posicao := uint(10)
	for i := uint(0); i <= posicao; i++ {
		fmt.Println(retornaPosicaoNumeroFibonnaci(i))
		time.Sleep(time.Second)
	}
}
