package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	slice := []int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println(slice)

	slice = append(slice, 80) //ADICIONANDO UM NOVO ITEM NO SLICE
	fmt.Println(slice)

	slice2 := array2[0:3]
	fmt.Println(slice2)

	array2[0] = "Posição alterada"
	fmt.Println(slice2)

	//Arrays internos
	fmt.Println("--------------------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade

}
