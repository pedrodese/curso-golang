package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo")
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
