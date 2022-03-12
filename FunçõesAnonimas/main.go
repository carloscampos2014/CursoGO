package main

import "fmt"

func main() {
	RetornaOperação("soma")(234, 5)
	RetornaOperação("subtração")(234, 5)
	RetornaOperação("multiplicação")(234, 5)
	RetornaOperação("divisão")(234, 5)
	RetornaOperação("bla")(234, 5)
}

func RetornaOperação(tipo string) func(n1, n2 int) {
	switch tipo {
	case "soma":
		return func(n1, n2 int) {
			fmt.Println(n1, "+", n2, "=", n1+n2)
		}
	case "subtração":
		return func(n1, n2 int) {
			fmt.Println(n1, "-", n2, "=", n1-n2)
		}
	case "multiplicação":
		return func(n1, n2 int) {
			fmt.Println(n1, "x", n2, "=", n1*n2)
		}
	case "divisão":
		return func(n1, n2 int) {
			fmt.Println(n1, "÷", n2, "=", n1/n2)
		}
	default:
		return func(n1, n2 int) {
			fmt.Println("Você passou uma operação entre os números", n1, "e", n2, "inexistente.")
		}
	}
}
