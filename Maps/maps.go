package main

import "fmt"

func main() {

	//O map não deixa usar dois tipos difentes, quando declaro um map, todos os atributos deverão ser do mesmo tipo
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Henrique",
	}

	//Forma de acessar apenas um atributo do map de usuario por exemplo.
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nomeCompleto": {
			"primeiroNome": "Pedro",
			"segundoNome":  "Henrique",
		},
		"faculdade": {
			"curso":    "Analise e Desenvolvimento de Sistemas",
			"semestre": "5 Semestre",
			"campus":   "UniSenai Joinville",
		},
	}
	fmt.Println(usuario2["nomeCompleto"])

	//forma de deletar uma chave de um map
	delete(usuario2, "faculdade")
	fmt.Println(usuario2)

	//forma de adicionar mais uma chave em um map já existente
	usuario2["signo"] = map[string]string{
		"signo": "Virgem",
	}

	fmt.Println(usuario2)
}
