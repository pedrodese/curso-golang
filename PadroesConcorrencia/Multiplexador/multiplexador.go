package main

import (
	"fmt"
	"time"
)

func main() {

	canalPrincipal := multiplexar(escrever("Olá Mundo!"), escrever("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canalPrincipal)
	}

}

// Padrão de projeto Multiplexador que consiste em receber mensagens de vários canais e repassá-las para um único canal.
func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <-canalEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()

	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}

	}()

	return canal
}
