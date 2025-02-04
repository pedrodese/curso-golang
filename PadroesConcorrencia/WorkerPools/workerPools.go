package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	//Usando 4 goroutines para processar as tarefas, seguindo o padrão de worker pools.
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- retornaPosicaoNumeroFibonnaci(numero)
	}
}

func retornaPosicaoNumeroFibonnaci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return retornaPosicaoNumeroFibonnaci(posicao-2) + retornaPosicaoNumeroFibonnaci(posicao-1)
}
