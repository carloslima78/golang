package main

import (
	"fmt"
	"os"
)

func main() {
	// Abre um arquivo
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}

	// Adia o fechamento do arquivo até o final da função main
	defer file.Close()

	// Escreve algo no arquivo
	_, err = file.WriteString("Hello, world!\n")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	// A função main está terminando aqui
	// O arquivo será fechado automaticamente antes do retorno
}
