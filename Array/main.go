package main

import (
	"fmt"
)

var lista []string

func main() {

	add("Carlos")
	add("Ana")
	fmt.Println(lista)
	add("Joana")
	add("Paula")
	fmt.Println(lista)
	delete(len(lista))
	fmt.Println(lista)
	delete(4)
}

func add(item string) {
	lista = append(lista, item)
}

func delete(indice int) {
	switch {
	case indice == 0:
		lista = lista[1:]
	case indice > 0 && indice < len(lista):
		lista = append(lista[:indice], lista[indice+1:]...)
	case indice == len(lista):
		lista = lista[:indice-1]
	default:
		fmt.Println("Elemento inexistente!")
	}
}
