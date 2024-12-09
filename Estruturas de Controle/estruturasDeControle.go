package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 20

	if numero > 10 {
		fmt.Println("O numero é maior do que 10")
	} else {
		fmt.Println("O numero é menor ou igual a 10")
	}

	//if init onde declaro uma variavel que irá ser acessivel apenas no escopo do if onde estou a declarando, fora desse escopo ela não existe.
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("O numero é maior do que zero")
	}
}
