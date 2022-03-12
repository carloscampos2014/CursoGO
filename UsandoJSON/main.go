package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome     string
	Idade    int
	Telefone string
}

func main() {
	lista := []Pessoa{}
	lista = append(lista, Pessoa{"Carlos Alberto", 45, "7866678"})
	lista = append(lista, Pessoa{"Jose Aparecido", 32, "4332423"})
	lista = append(lista, Pessoa{"Luana Morais", 12, "55353"})

	fmt.Println(lista)
	fmt.Println("Convertendo em JSON")
	j, err := json.Marshal(lista)
	if err != nil {
		fmt.Println("Erro:", err.Error())
	} else {
		fmt.Println(string(j))
	}
	fmt.Println("Convertendo em GO Struc")
  novaPessoa := []Pessoa{}
	err = json.Unmarshal(j, &novaPessoa)
	if err != nil {
		fmt.Println("Erro:", err.Error())
	} else {
		fmt.Println(novaPessoa)
	}

}
