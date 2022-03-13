package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := Converge(Trabalho("Montar Coisa"), Trabalho("Embalar Coisa"))

	for i := 0; i < 20; i++ {
		fmt.Printf("Ciclo %v:\n  Mensagem: %v\n", i, <-canal)
	}
}

func Trabalho(texto string) chan string {
	canal := make(chan string)

	go func(s string, c chan string) {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Função %v diz: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(texto, canal)

	return canal
}

func Converge(a, b chan string) chan string {
	novo := make(chan string)
	go func() {
		for v := range a {
			novo <- v
		}
	}()

	go func() {
		for v := range b {
			novo <- v
		}
	}()
	return novo
}
