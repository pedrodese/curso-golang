package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"

	retorno := func() {
		fmt.Println(texto)
	}

	return retorno
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
