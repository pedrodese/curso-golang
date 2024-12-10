package main

import "fmt"

func main() {

	//a função anonima pode ser usada tanto dessa forma, como da forma abaixo, podendo ou não conter um retorno ou um parametro por exemplo
	func(texto string) {
		fmt.Println(texto)
	}("Essa é uma função anonima")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebendo parametro e armazenando valor da função anonima em uma variavel -> %s ", texto)
	}("Parametro recebido")

	fmt.Println(retorno)
}
