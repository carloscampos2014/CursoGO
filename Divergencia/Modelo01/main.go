package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	canal01 := make(chan int)
	canal02 := make(chan int)

	go mandar(90000, canal01)

	go recebe(canal01, canal02)

	for v := range canal02 {
		fmt.Printf("%v\t", v)
	}
}

func mandar(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
	}
	close(c)
}

func recebe(a, b chan int) {
	var wg sync.WaitGroup
	for v := range a {
		wg.Add(1)
		go func(x int) {
			b <- trabalho(x)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(b)
}

func trabalho(i int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(3e0)))
	return i * rand.Intn(30)
}
