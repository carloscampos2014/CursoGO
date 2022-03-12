package main

import "fmt"

func main() {
	fmt.Println("O Fatorial de 5 é ", Fatorial(5))
	fmt.Println("O Fatorial de 32 é ", Fatorial(32))
	fmt.Println("O Fatorial de 3456 é ", Fatorial(3456))
}

func Fatorial(x float64) float64 {
	if x == 0 || x == 1 {
		return 1
	}
	return x * Fatorial(x - 1)
}
