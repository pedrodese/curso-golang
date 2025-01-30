package main

import "fmt"

func main() {
	//O canal com buffer é um canal que pode armazenar um número limitado de valores, ou seja, ele pode armazenar valores até que ele esteja cheio. caso o canal esteja cheio, a goroutine que tentar enviar um valor para o canal ficará bloqueada até que um valor seja retirado do canal
	canal := make(chan string, 2)

	canal <- "Olá Mundo"

	mensagem := <-canal
	fmt.Println(mensagem)
}
