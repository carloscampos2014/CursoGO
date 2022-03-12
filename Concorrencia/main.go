package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main(){

	fmt.Printf("Numero de Processadores: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de Rotinas: %v\n", runtime.NumGoroutine())

	wg.Add(2)

	go func1()
	go func2()

	fmt.Printf("Numero de Rotinas: %v\n", runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("Numero de Rotinas: %v\n", runtime.NumGoroutine())
}

func func1(){
	for i := 0; i <= 10; i++ {
		fmt.Printf("Função 01: %v\n", i)		
		time.Sleep(200)
	}
	wg.Done()
}

func func2(){
	for i := 0; i <= 10; i++ {
		fmt.Printf("Função 02: %v\n", i)		
		time.Sleep(200)
	}
	wg.Done()
}