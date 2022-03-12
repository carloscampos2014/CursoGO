package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
)

const mensagemErro = `
Sintaxe do comando Errada!
Sintaxe Criptografia <função>
Argumentos:
  <função> - A Função que vai ser utilizada
      gerar - Para gerar um Hash novo
      pegar - Usar hash pre compilado

`
var sBytes []byte 

var err error

func main() {
	toGetAllArgs := os.Args[1:]
	senha := "05031996"
	if len(toGetAllArgs) < 1 {
		fmt.Print(mensagemErro)
		return
	}

	switch toGetAllArgs[0] {
	case "gerar":
		sBytes, err = GeraHash(senha)
	case "pegar":
		sBytes, err = PegarHash("$2a$10$TJbWlU.z/FNewzYXoW00jeVk/WPyhbKfqkNNftaFpGJYdWZpYH7uq")
	default:
		fmt.Print(mensagemErro)
		return
	}

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Senha:", senha, "Hash:", string(sBytes))
	err = bcrypt.CompareHashAndPassword(sBytes, []byte("270210"))
	if err != nil {
		fmt.Println("Senha Incorreta!")
	} else {
		fmt.Println("Senha Correta!")
	}
	
	err = bcrypt.CompareHashAndPassword(sBytes, []byte("05031996"))
	if err != nil {
		fmt.Println("Senha Incorreta!")
	} else {
		fmt.Println("Senha Correta!")
	}
}

func GeraHash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), 10)
}

func PegarHash(texto string) ([]byte, error) {
	return []byte(texto), nil
}
