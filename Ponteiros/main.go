package main

import "fmt"

type pessoa struct {
	nome     string
	telefone string
}

func main() {
	pessoa1 := pessoa{"José", "7867868778"}
	fmt.Println(pessoa1)
	MudeNome(&pessoa1, "Carlos")
	fmt.Println(pessoa1)
	MudeTelefone(&pessoa1, "1121131")
	fmt.Println(pessoa1)
}

func MudeNome(p *pessoa, novo string) {
	(*p).nome = novo
}

func MudeTelefone(p *pessoa, novo string) {
	(*p).telefone = novo
}
