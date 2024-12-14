package main

import "fmt"

func main() {
	fmt.Println(alunoEstaAprovado(7, 9))
}

//defer = ADIAR, isso significa que posso adiar uma rotina do meu programa, sendo muito util em diversas ocasiões, o defer em funções que contem retorno, ira ser executado
//antes de qualquer return na função abaixo isso está exemplificado, onde ele ira printar a mensagem "Média calculado. Resultado será retornado!" em qualquer return que ele der
//mesmo sendo true ou false
func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Entrando na função que verifica se o aluno está aprovado ")
	media := (n1 + n2) / 2

	if media >= 7 {
		return true
	}
	return false

}
