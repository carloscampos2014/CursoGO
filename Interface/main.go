package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) DizerBomDia() {
	fmt.Printf("Bom dia, meu nome é %v, Tudo bem com Você?\n", p.nome)
}

type dentista struct {
	pessoa
	crm      string
	consulta float64
}

type arquiteto struct {
	pessoa
	crea     string
	projetos string
	salario  float64
}

type cabelereiro struct {
	pessoa
	corteMasculino float64
	corteFeminino  float64
}

func (d dentista) DizerDadosProfissionais() {
	fmt.Printf("Meu CRM é %v, eu cobro R$ %9.2f por consulta.\n", d.crm, d.consulta)
}

func (a arquiteto) DizerDadosProfissionais() {
	fmt.Printf("Meu CREA é %v, eu cobro no mínimo R$ %9.2f por projeto.\n Eu faço esse tipo de projetos: %v.\n", a.crea, a.salario, a.projetos)
}

func (c cabelereiro) DizerDadosProfissionais() {
	fmt.Printf("Meus cortes de cabelo custam:\n  Feminino  R$ %9.2f\n  Masculino R$ %9.2f\n", c.corteFeminino, c.corteMasculino)
}

type IResposta interface {
	DizerBomDia()
	DizerDadosProfissionais()
}

func DizerResposta(resposta IResposta) {
	resposta.DizerBomDia()
	resposta.DizerDadosProfissionais()
}

func main() {
	arq := arquiteto{
		pessoa{"Joana Prado", 31},
		"94565",
		"Cozinhas, Salas e Banheiros",
		350.00,
	}

	dent := dentista{
		pessoa{"Paulo Augusto", 29},
		"13245",
		56.78,
	}

	cab := cabelereiro{
		pessoa:         pessoa{"Patricia Alessandro", 37},
		corteFeminino:  54.00,
		corteMasculino: 25.00,
	}

	fmt.Printf("\n\n")
	DizerResposta(arq)
	fmt.Printf("\n\n")
	DizerResposta(dent)
	fmt.Printf("\n\n")
	DizerResposta(cab)
	fmt.Printf("\n\n")
}
