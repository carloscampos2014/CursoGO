package main

import (
	"fmt"
	"sort"
)

type Pessoas struct {
	Nome  string
	Idade int
}

type OrdenarPorNome []Pessoas

func (o OrdenarPorNome) Len() int           { return len(o) }
func (o OrdenarPorNome) Less(i, j int) bool { return o[i].Nome < o[j].Nome }
func (o OrdenarPorNome) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }

type OrdenarPorIdade []Pessoas

func (o OrdenarPorIdade) Len() int           { return len(o) }
func (o OrdenarPorIdade) Less(i, j int) bool { return o[i].Idade < o[j].Idade }
func (o OrdenarPorIdade) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }

func main() {
	fmt.Println("Ordenando Textos")
	
	frutas := []string{"Morango", "Abacaxi", "Pera", "Banana", "Uva", "Coco", "Abacate"}
	fmt.Println("  ", frutas)
	
	sort.Strings(frutas)
	fmt.Println("  ", frutas)
	
	sort.Sort(sort.Reverse(sort.StringSlice(frutas)))
	fmt.Println("  ", frutas)
	
	fmt.Println("Ordenando Números")
	
	números := []int{5, 23, 67, 2, 15, 7, 89, 1, 69, 32}
	fmt.Println("  ", números)
	
	sort.Ints(números)
	fmt.Println("  ", números)
	
	sort.Sort(sort.Reverse(sort.IntSlice(números)))
	fmt.Println("  ", números)
	
	fmt.Println("Ordenando Customizada")

	pessoa := []Pessoas{}
	pessoa = append(pessoa, Pessoas{"Janaina", 30})
	pessoa = append(pessoa, Pessoas{"Carlos", 22})
	pessoa = append(pessoa, Pessoas{"Renan", 5})
	pessoa = append(pessoa, Pessoas{"Adriana", 70})

	fmt.Println("Original                     ->", pessoa)

	sort.Sort(OrdenarPorNome(pessoa))
	fmt.Println("Ordenado por Nome            ->", pessoa)

	sort.Sort(sort.Reverse(OrdenarPorNome(pessoa)))
	fmt.Println("Ordenado Invertido por Nome  ->", pessoa)

	sort.Sort(OrdenarPorIdade(pessoa))
	fmt.Println("Ordenado por Idade           ->", pessoa)

	sort.Sort(sort.Reverse(OrdenarPorIdade(pessoa)))
	fmt.Println("Ordenado Invertido por Idade ->", pessoa)
}
