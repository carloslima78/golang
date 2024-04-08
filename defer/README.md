# Defer

Em Go, `defer` é uma declaração que é usada para adiar a execução de uma função até que a função envolvente (a função que contém o defer) retorne. Isso pode ser útil para garantir que determinadas operações sejam executadas antes que a função retorne, como fechar um arquivo, liberar um recurso, ou executar limpezas necessárias.

Quando uma função é chamada com `defer`, o argumento da função é avaliado imediatamente, mas a execução da função em si é adiada até que a função envolvente complete. Isso é particularmente útil para garantir que operações de limpeza ou fechamento sejam sempre executadas, independentemente de como a função é encerrada (por exemplo, retornando normalmente ou gerando um erro).

Aqui está um exemplo simples:

```go
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
```

Neste exemplo, `defer file.Close()` garante que o arquivo seja fechado automaticamente antes que a função `main` retorne. Mesmo que ocorra um erro ao escrever no arquivo, ou mesmo se o programa for encerrado antes do final, o arquivo será fechado de forma adequada, evitando vazamentos de recursos.