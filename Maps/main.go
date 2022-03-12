package main

import "fmt"

type agenda struct {
	nome     string
	telefone string
}

func main() {
	lista := map[int]agenda{}

	lista[0] = agenda{nome: "Carlos", telefone: "(11) 99674-7477"}
	lista[1] = agenda{nome: "Consuelo", telefone: "(11) 96857-8924"}
	lista[2] = agenda{nome: "Gabriel", telefone: "(11) 98674-7477"}

	fmt.Println(lista)

	lista[0] = agenda{nome: "Carlos Alberto", telefone: lista[0].telefone}

	fmt.Println(lista[0].nome)

	fmt.Println(lista)

}
