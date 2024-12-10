package main

import "fmt"

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("PÓS EXECUÇÃO DA FUNÇÃO")
}

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	//o panic interrompe completamente minha aplicação, nesse caso quando a média for exatamente 6 ele interrompe completamente a execução
	panic("A MEDIA É EXATAMENTE 6")
}
