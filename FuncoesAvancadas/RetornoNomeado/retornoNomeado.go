package main

import "fmt"

func main() {
	fmt.Println("RETORNOS NOMEADOS")
	soma, subtracao := calculoMatematico(20, 20)
	fmt.Println(soma, subtracao)
}

//Nessa função tenho um retorno nomeado, dessa forma não preciso inicializar a variavel dentro da função e meu return é global, pois a função já sabe oq vai ser retornado.
func calculoMatematico(n1 int, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}
