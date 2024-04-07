# Módulos

Em Go, um módulo é uma unidade de organização de dependências. Cada módulo é um conjunto de pacotes que são versionados e gerenciados em conjunto pelo sistema de módulos do Go. 

Observa-se que em Go, um módulo não é exatamente o mesmo conceito que uma classe ou pacote em outras linguagens como C# ou Java. 


## Entendendo a estrutura de um módulo

Se temos uma aplicação simples com funcionalidades relacionadas e cliente e produto, é provável que eles seriam parte do mesmo módulo, a menos que você tenha uma razão específica para dividir em módulos separados.

```golang
meuapp/
├── go.mod
└── pkg
    ├── cliente
    │   └── cliente.go
    ├── produto
    │   └── produto.go
    └── ...
```

Neste exemplo, `meuapp` é o nome do módulo e dentro dele você tem os pacotes `cliente` e `produto`. O arquivo `go.mod` no diretório raiz do módulo (`meuapp`) define as dependências e versões para todo o módulo.

Se `cliente` e `produto` fossem módulos separados, você teria uma estrutura mais ou menos assim:

```golang
meuapp/
├── go.mod
├── cliente/
│   ├── go.mod
│   └── cliente.go
└── produto/
    ├── go.mod
    └── produto.go
```

Neste caso, `cliente` e `produto` seriam módulos separados com seus próprios arquivos `go.mod` dentro de seus diretórios.

Em resumo, em Go, podemos ter um módulo que contém vários pacotes relacionados, ou pode ter vários módulos se quiser uma separação mais rígida entre diferentes partes do seu código. A escolha depende da estrutura e da complexidade do seu projeto.


## Aplicação de exemplo

Vamos criar um exemplo simples onde temos um pacote `cliente` que importa e instancia um pacote `produto`.

1. Primeiro, vamos criar a estrutura de diretórios e arquivos:

```golang
meuapp/
└── pkg
    ├── cliente
    │   └── cliente.go
    ├── produto
    │   └── produto.go
    └── ...
```

2. Dentro do diretório meuapp, inicialize o módulo Go com o comando:

```bash
go mod init meuapp
```

Após a execução do comando acima, a espera-se que tenha sido criado o arquivo `go.mod` na estrutura de pastas:

```golang
meuapp/
├── go.mod
├── cliente/
│   ├── cliente.go
└── produto/
    └── produto.go
```

3. Agora, vamos criar o conteúdo dos arquivos:

Dentro de `produto/produto.go`:

```golang
package produto

// Produto representa um produto
type Produto struct {
    Nome string
}

// NovoProduto cria uma nova instância de Produto
func NovoProduto(nome string) *Produto {
    return &Produto{
        Nome: nome,
    }
}
```

Dentro de `cliente/cliente.go`:

```golang
package cliente

import (
	"fmt"
	"meuapp/produto"
)

// Cliente representa um cliente
type Cliente struct {
	Nome    string
	Produto produto.Produto // Instancia um Produto
}

// NovoCliente cria uma nova instância de Cliente
func NovoCliente(nome string, produto produto.Produto) *Cliente {
	return &Cliente{
		Nome:    nome,
		Produto: produto,
	}
}

// ExibirDados exibe os dados do cliente e do produto
func (c *Cliente) ExibirDados() {
	fmt.Printf("Cliente: %s\n", c.Nome)
	fmt.Printf("Produto: %s\n", c.Produto.Nome)
}
```

4. Agora, vamos criar um arquivo main.go no diretório meuapp para utilizar esses pacotes:

```golang
package main

import (
    "meuapp/cliente"
    "meuapp/produto"
)

func main() {
    // Criando um novo produto
    meuProduto := produto.NovoProduto("Computador")

    // Criando um novo cliente com o produto
    meuCliente := cliente.NovoCliente("Alice", *meuProduto)

    // Exibindo os dados do cliente e do produto
    meuCliente.ExibirDados()
}
```

5. Agora, para executar o exemplo, basta rodar o seguinte comando no diretório meuapp:

```bash
go run main.go
```

Isso vai criar um novo produto "Computador" e associá-lo ao cliente "Alice". Em seguida, ele exibirá os dados do cliente e do produto:

```bash
Cliente: Alice
Produto: Computador
```

Neste exemplo, o pacote cliente importa o pacote produto e consegue instanciar um Produto dentro da struct Cliente. Isso ilustra como um pacote pode usar outro pacote do mesmo módulo em Go.
