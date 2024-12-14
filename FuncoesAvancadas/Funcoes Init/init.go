package main

import "fmt"

var numero int

func init() {
	fmt.Println("Rodando função init")
	numero = 20
}

func main() {
	fmt.Println("Rodando função main")
	fmt.Println(numero)
}
