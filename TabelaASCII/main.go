package main

import "fmt"

func main() {
	posicao := 1
	for letra := 32; letra <= 122; letra++ {
		fmt.Printf("%#c\t", letra)
		if posicao > 11 {
			posicao = 0
			fmt.Println("")
		}
		posicao++
	}
}
