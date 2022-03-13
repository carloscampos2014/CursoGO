package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)

	go Enviar(par, impar)
	go Recebe(par, impar, converge)

	for v := range converge {
		fmt.Println("Numero de Convergência ->", v)
	}
}

func Enviar(par, impar chan int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
}

func Recebe(par, impar, converge chan int) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for v := range par {
			converge <- v
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for v := range impar {
			converge <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(converge)
}
