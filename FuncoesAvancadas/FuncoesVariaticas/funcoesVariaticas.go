package main

import "fmt"

func main() {

	resultado := somaVariosNumeros(10, 10, 10, 10, 10, 10, 10)
	fmt.Println(resultado)

}

// nessa função variatica, não preciso especificar a quantidade de parametros que irei receber, pois ele transforma em um slice, onde posso percorrer com um for por exemplo
// nas funções variaticas, posso receber apenas um parametro do tipo variatico, e por padrão ele dever ser o ultimo parametro a ser recebido na função
func somaVariosNumeros(numeros ...int) int {
	total := 0

	for _, somaDosValores := range numeros {
		total += somaDosValores
	}

	return total
}
