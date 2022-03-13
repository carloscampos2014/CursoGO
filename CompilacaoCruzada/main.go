package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Aplicativo Compilação Cruzada\n  Gerado num Sistema\t-> windows/amd64\n  Executado num Sistema\t-> %v/%v\n", runtime.GOOS, runtime.GOARCH)
}
