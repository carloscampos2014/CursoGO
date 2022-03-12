package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		dado := (rand.Int31n(7)) + 1
		if dado < 7 {
			fmt.Println("Dessa vez o dado caiu numero ->", dado)
		}
		switch dado {
		case 1:
			fmt.Println("Vai ser azarado assim na china!")
		case 2:
			fmt.Println("Ta melhor!")
		case 3:
			fmt.Println("Se tivesse mais sorte!")
		case 4:
			fmt.Println("Quase lá")
		case 5:
			fmt.Println("Passou Raspando!")
		case 6:
			fmt.Println("Vai ser sortudo assim na china!")
		}
		if dado == 6 {
			break
		}
	}
}
