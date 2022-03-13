package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go Envia(par, impar, quit)

	Recebe(par, impar, quit)
}

func Envia(par, impar chan int, quit chan bool) {
	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
	quit <- true
}

func Recebe(par, impar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Printf("O numero %v é par.\n", v)
		case v := <-impar:
			fmt.Printf("O numero %v é impar.\n", v)
		case v, ok := <-quit:
			if ok {
				fmt.Println("Finalizado ->", v)
				return
			} else {
				fmt.Println("Finalizado ->", v)
			}

		}

	}
}
