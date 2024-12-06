package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Pedro"
	u.idade = 21
	fmt.Println(u)

	endereco := endereco{"Rua dos Bobos", 0}

	usuario2 := usuario{"Pedro", 21, endereco}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Pedro"}
	fmt.Println(usuario3)
}
