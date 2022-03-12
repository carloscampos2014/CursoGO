package main

import "fmt"

func main() {
	lista := []string{}
	fmt.Println(lista)
	Acrecentar(&lista,"bla")
	fmt.Println(lista)

	for i := 0; i < 256; i++ {
		Acrecentar(&lista,string(i))
	}
	fmt.Println(lista)

}

func Acrecentar(l *[]string, texto string) {
	*l = append(*l, texto)
}
