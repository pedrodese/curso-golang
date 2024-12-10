package main

import (
	"fmt"
	"time"
)

// lembrando que o Go só tem apenas um laço de repetição que é o FOR que pode ser usado de diferentes formas
func main() {
	fmt.Println("Loops de repetição")

	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando I")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando J ", j)
		time.Sleep(time.Second)
	}

	//forma de iterar algo, parecido com o foreach do Java
	nomes := []string{"Pedro", "Rafaela", "Kelcy"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	//posso tambem iterar uma string percorrendo ela e retornando os indices e os valores
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Pedro Henrique",
		"profissao": "Desenvolvedor Back-End",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
