package main

import "fmt"

func main() {
	//forma de declarar variaveis de forma explicita
	var string1 string = "string1"

	//forma de declarar variaveis de forma implicita
	string2 := "string2"

	fmt.Println(string1)
	fmt.Println(string2)

	//outra forma de declarar variaveis de forma explicita
	var (
		string3 string = "lalala"
		string4 string = "lalala"
	)
	fmt.Println(string3, string4)

	string5, string6 := "string5", "string6"
	fmt.Println(string5, string6)
}
