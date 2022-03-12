package main

import "fmt"

func main() {
	for ano := 2020; ano < 2051; ano++ {
		texto := "é Bissexto!"
		if ano%4 == 0 {
			texto = "não é Bissexto!"
		}
		fmt.Printf("Ano %v %v\n", ano, texto)
		for mes := 1; mes <= 12; mes++ {
			fmt.Printf("%v\t", mes)
		}
		fmt.Println("")
	}
}
