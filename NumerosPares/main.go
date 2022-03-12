package main

import "fmt"

func main() {
	for numero := 0; numero < 21; numero++ {
		if numero%2 == 0 {
			fmt.Println("Numero", numero, "é um numero   par!")
		} else {
			fmt.Println("Numero", numero, "é um numero ímpar!")
		}
	}
}
