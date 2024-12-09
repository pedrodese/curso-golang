package main

import "fmt"

func retornaDiaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabado"
	default:
		return "Numero Inválido"
	}
}

//Segundo jeito de se utilizar o switch onde faço a comparação dentro do proprio case, utilizando a variavel que recebemos na função
func retornaDiaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sabado"
	default:
		return "Numero invalido"
	}
}

func main() {
	fmt.Println("Switchs")

	diaDaSemana := retornaDiaDaSemana(6)
	diaDaSemana2 := retornaDiaDaSemana2(7)

	fmt.Println(diaDaSemana)
	fmt.Println(diaDaSemana2)
}
