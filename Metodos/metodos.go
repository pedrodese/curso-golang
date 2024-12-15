package main

import "fmt"

//metodos estão obrigatoriamente vinculados a alguma estrutura, nesse caso todos os metodos estão vinculados ao struct de usuario.
type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Salvando informações do Usuario: %s \n", u.name)
}

func (u user) validateLegalAge() bool {
	return u.age >= 18
}

func (u *user) makeBirthday() {
	u.age++
}

func main() {

	firstUser := user{"Pedro", 21}
	firstUser.save()

	secondUser := user{"Rafaela", 23}
	secondUser.save()

	userIsLegalAge := secondUser.validateLegalAge()
	fmt.Println(userIsLegalAge)

	firstUser.makeBirthday()
	fmt.Println(firstUser.age)

}
