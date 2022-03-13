package main

import (
	"fmt"
)

func main() {
	números := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	canal := make(chan int)
	go Send(canal, números...)
	Receive(canal)
}

func Send(s chan<- int, num ...int) {
	for _, v := range num {
		s <- v * 10
	}
	close(s)
}

func Receive(r <-chan int) {
	for v := range r {
		fmt.Println("Valor Recebido ->", v)
	}
}
