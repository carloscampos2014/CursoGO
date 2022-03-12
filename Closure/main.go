package main

import "fmt"

func main() {
	amigos, família := MontaLista(), MontaLista()

	fmt.Println("Lista Amigos")
	amigos("Sakura")
	amigos("Naruto")
	amigos("Sasuke")
	amigos("Kakashi")
	fmt.Println("Lista Família")
	família("Minato")
	família("Kushina")
	família("Naruto")
}

func MontaLista() func(nome string) {
	lista := []string{}
	return func(nome string) {
		lista = append(lista, nome)
		fmt.Println("Dados da Lista:", lista)
	}
}
