package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       uint8
}

type student struct {
	person
	course   string
	semester uint8
}

func main() {
	fmt.Println("Herança só que não")

	//CRIANDO PESSOA
	pessoa := person{"Pedro", "Henrique", 21}

	//CRIANDO ESTUDANTE
	estudante := student{pessoa, "Análise e desenvolvimento de sistemas", 4}
	fmt.Println(estudante)

	//PARA ACESSAR UM ATRIBUTO DE UM STRUCT QUE ESTÁ HERDANDO OUTRO STRUCT FAZEMOS DESSA FORMA
	primeiroNomeEstudante := estudante.firstname
	fmt.Println(primeiroNomeEstudante)

}
