## Tracing no Go: pprof

Trata-se de uma ferramenta de profiling e tracing para aplicações Go. Ela oferece várias funcionalidades, incluindo:

- Profiling da CPU para identificar onde o programa gasta mais tempo.
- Profiling de alocação de memória para descobrir onde e como a memória está sendo alocada.
- Trace profiling para entender o fluxo de execução do programa e quais funções estão sendo chamadas.

### Como utilizar o ppprof no Go?

- Importe o pacote "net/http/pprof" no código:

```golang
import _ "net/http/pprof"
```

- Adicione uma rota no seu servidor HTTP para acessar os perfis do pprof:

```golang
func main() {
    // Outro código do seu servidor...

    // Rota para perfis do pprof
    http.HandleFunc("/debug/pprof/", pprof.Index)
    http.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
    http.HandleFunc("/debug/pprof/profile", pprof.Profile)
    http.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
    http.HandleFunc("/debug/pprof/trace", pprof.Trace)

    // Inicie o servidor HTTP
    http.ListenAndServe(":8080", nil)
}
```

- Agora, quando o seu servidor estiver rodando, você pode acessar os perfis do pprof em seu navegador. Por exemplo:

- http://localhost:8080/debug/pprof/ para ver uma lista de perfis disponíveis.
- http://localhost:8080/debug/pprof/profile para gerar um perfil de CPU.
- http://localhost:8080/debug/pprof/heap para gerar um perfil de alocação de memória.

É possível usar ferramentas como go tool pprof para analisar esses perfis. Por exemplo:

```sh
go tool pprof http://localhost:8080/debug/pprof/profile
```

Isso abrirá uma interface de linha de comando interativa onde pode-se executar comandos para ver as funções que estão consumindo mais tempo de CPU ou memória.

O **pprof** é uma ferramenta essencial para qualquer desenvolvedor Go que queira entender e otimizar o desempenho de suas aplicações. 

Ele oferece muito mais recursos avançados para análise de desempenho e diagnóstico de problemas em aplicações Go. É uma ferramenta poderosa para entender o comportamento de uma aplicação em termos de consumo de CPU, alocação de memória e muito mais.

## Exemplo de uso

Vamos criar um código simples que gere um gargalo de CPU. Para isso, criaremos uma função que realiza um cálculo repetitivo e intensivo em termos de CPU. Em seguida, utilizaremos o pprof para analisar esse gargalo.

```golang
package main

import (
	"fmt"
	"math/rand"
	"time"

	_ "net/http/pprof"
)

func main() {
	// Importando o pacote net/http/pprof para habilitar os perfis do pprof
	go func() {
		fmt.Println("Starting pprof server on :6060")
		err := http.ListenAndServe(":6060", nil)
		if err != nil {
			fmt.Println("Error starting pprof server:", err)
		}
	}()

	for i := 0; i < 10; i++ {
		go calculation()
	}

	// Para manter o programa rodando para coletar o perfil do pprof
	fmt.Println("Pressione Ctrl+C para encerrar.")
	select {}
}

func calculation() {
	for {
		result := 0
		for i := 0; i < 100000; i++ {
			// Realizando algum cálculo intensivo de CPU
			result += rand.Intn(1000)
		}
		fmt.Println("Result:", result)
		time.Sleep(100 * time.Millisecond)
	}
}
```

Neste exemplo, a função calculation() realiza um cálculo intensivo de CPU repetitivo. Ela gera um número aleatório entre 0 e 1000 e adiciona ao resultado, repetindo isso várias vezes. Isso cria um gargalo de CPU.

## Como usar o pprof para analisar o gargalo de CPU:

- Compile e execute o código acima. O servidor pprof será iniciado em localhost:6060.

```golang
go run main.go
```

- Em outro terminal, use a ferramenta go tool pprof para acessar o perfil de CPU:

```bash
go tool pprof http://localhost:6060/debug/pprof/profile
```

Isso abrirá uma interface de linha de comando interativa. Aguarde um pouco enquanto o perfil é coletado.

- Uma vez que o perfil estiver carregado, você verá um prompt (pprof). Digite top para ver as funções que estão consumindo mais tempo de CPU:

```bash
(pprof) top
```

Isso mostrará uma lista das funções com o maior tempo de CPU. Procure pela função main.calculation, que é a nossa função com o cálculo intensivo de CPU. A coluna % mostra o percentual do tempo de CPU gasto em cada função.

- Para obter mais informações sobre uma função específica, digite list seguido do nome da função:

```bash
(pprof) list calculation
```

Isso mostrará o trecho de código onde o gargalo está ocorrendo, com indicações das linhas:

```bash
Total: 0.5s
ROUTINE ======================== main.calculation in /path/to/your/file/main.go
  400ms   80.49%   80.49%      400ms   main.calculation
  100ms   19.51%  100.00%      500ms   main.main.func1
    0ms    0.00%  100.00%      500ms   runtime.goexit
```

Isso mostra que a maior parte do tempo é gasta na função main.calculation.

Trata-se de um exemplo simples de como usar o pprof para identificar e analisar gargalos de CPU em um programa Go. Com essa ferramenta, é possível identificar quais partes do código estão consumindo mais recursos e otimizá-las conforme necessário.