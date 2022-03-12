package main

import "fmt"

func main() {
	números := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("A soma dos números %v é %v\n", números, soma(números...))
	fmt.Printf("A soma somente dos pares dos números %v é %v\n", números, SomaNúmerosPares(soma,números...))
	fmt.Printf("A soma somente dos ímpares dos números %v é %v\n", números, SomaNúmerosÍmPares(soma,números...))
}

func soma(num ...int) int {
	total := 0
	for _, valor := range num {
		total += valor
	}
	return total
}

func SomaNúmerosPares(f func(num ...int) int, números ...int) int {
	pares := []int{}
	for _, valor := range números {
		if valor%2 == 0 {
			pares = append(pares, valor)
		}
	}
	total := f(pares...)
	return total
}

func SomaNúmerosÍmPares(f func(num ...int) int, números ...int) int {
	pares := []int{}
	for _, valor := range números {
		if valor%2 != 0 {
			pares = append(pares, valor)
		}
	}
	total := f(pares...)
	return total
}
