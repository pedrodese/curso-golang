package main

import (
	"errors"
	"fmt"
)

func main() {

	//inteiro
	var numero int64 = -20000000000000
	fmt.Println(numero)
	var numero2 uint32 = 20000
	fmt.Println(numero2)
	//existem varios alias - (apelidos) para os tipos de dados
	var numero3 rune = 12345
	fmt.Println(numero3)
	var numero4 byte = 123
	fmt.Println(numero4)

	//numeros reais
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)
	var numeroReal2 float64 = 123456789.000
	fmt.Println(numeroReal2)

	// strings
	var str string = "Pedro"
	fmt.Println(str)
	str2 := "Pedro 2"
	fmt.Println(str2)

	// caractere unico ele deve ser entre aspas simples e retornara um numero baseado na tabela ask
	char := 'F'
	fmt.Println(char)

	//fim strings

	//boleanos

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)

	//criando um erro personalizado
	var erro2 error = errors.New("ERRO INTERNO")
	fmt.Println(erro2)

}
