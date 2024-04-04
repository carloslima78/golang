# Golang (Go)

Bem-vindos ao estudo da linguagem de programação Go (Golang)! Go é uma linguagem de programação moderna, eficiente e de código aberto criada pelo Google. Com uma sintaxe simples e poderosa, Go é projetada para facilitar o desenvolvimento de software escalável e concorrente.

## Forças

- Sintexe simples de aprender, curva de aprendizado menor que outras linguagens.
- Oferece performance, concorrência e paralelismo.
- Reduz o consumo de poder computacional, disponibliza mecanismos que colobora exigir menos recursos.
- Plataforma com recursos nativos de segurança.
- Bastante utilizado para comunicação entre microservices via GRPC.
- Acesso simples aos recursos e protocolos modernos de rede.
- Nasceu otimizado para multi-core.
- Baixa utilização de memória, inicia uma thread utilizando apenas 2k.
- Garbage Collector otimizado.
- O Go é uma linguagem compilada, porém não utiliza máquina virtual.


## Timeline

- **2007:** Início do projeto.
- **2009:** Open Sourced.
- **2012:** Versão estável v1.0.
- **2015:** Melhorias importantes v1.5.
- **2018:** Consolidação do ecossistema v1.11.
- **2022:** Introdução do Generics v1.18.
- **2024:** Novidades no HTTP Server para gerenciamento de rotas v1.22.

## Retrocompatilidade

O Golang oferece retrocompatiiblidade com todas as versões anteriores, ou seja, uma versão mais atualualizada do Go, vai conseguir executar aplicações compiladas em qualquer versão anterior independente da evolução e do tempo.

## Go é orientado a objetos?

O Go é uma linguagem fortemente tipada, porém, *não é orientada a objetos* e sim, orientada a dados.

## Go é Plataforma

- Programação.
- Testes
- Geração de documentos.
- Gerenciamento de dependências.
- Multiplataforma.
- Profilling.
- Security.
- Performance "out-of-box"
- Build & Deploy.


## Instalação

- Acesse  *go.dev*, selecione o sistema operacional e sucesso.
- Instale o plugin oficial para VSCode *Rich Go language support for Visual Studio Code*.


## Packages

Packages são como o Go organiza o código em um programa. Cada programa Go precisa de um pacote de entrada único. Um pacote Go não pode conter dois pacotes com nomes diferentes. 

Geralmente, um programa Go começa pela função main, que é a entrada do programa. A partir da função main, é possível chamar outras funções ou pacotes para executar diversas tarefas no programa.

[Link para pkg.go.dev que contém os pacotes padrão do Go](https://pkg.go.dev/std)

1. Crie um arquivo *main.go* e escreva o código abaixo:

```golang
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Olá Mundo")
}
```

2. Comando para executar um progrma Go

```golang
go run main.go
```


## Visibilidade

No Go, não existem as visibilidades presentes no encapsulamento da orientação a objetos, a visibilidade é controlada pelo uso de letras maiúsculas ou minúsculas no início dos nomes. 

Variáveis ou funções que começam com letra maiúscula são visíveis (exportadas), podendo ser acessadas por pacotes externos. Por exemplo, a variável Nome é exportada, enquanto a função calcularMedia não é. Veja:

```golang
package main

import (
	"fmt"
)

// Variável exportada
var Nome string = "João"

// Função não exportada
func calcularMedia(n1, n2 float64) float64 {
	return (n1 + n2) / 2
}

func main() {
	fmt.Println(Nome) // Saída: João
	// fmt.Println(calcularMedia(7.5, 8.2)) // Erro, função não exportada
}
```
Na função main, podemos imprimir a variável Nome, pois ela é exportada e visível em outros pacotes. No entanto, ao tentar chamar calcularMedia fora deste pacote, ocorrerá um erro, pois ela não é exportada e, portanto, não é visível em outros pacotes.


## Variáveis e Tipos de Dados

Variáveis são usadas para armazenar dados em um programa Go, já os tipos de dados são usados para especificar o tipo de dado que uma variável pode armazenar.

No Go, todo tipo de dado é fortemente tipado, e isso significa que cada variável deve ser declarada com um tipo de dado específico. 

Toda variável no Go é iniciada com um valor padrão, exemplo:

- int = 0
- string = ""
- bool = false
- float64 = 0.0


### Variável a nível de Package

Uma variável a nível de package é visível para todas as funções do pacote. Para declarar uma variável a nível de package, use a palavra-chave `var` seguida do tipo de dado e do nome da variável.

Por exemplo, a seguinte declaração cria uma variável a nível de package chamada `name` do tipo `string`:

```golang
package main

import (
	"fmt"
)

// declaração da variável name, que armazenará um texto
var name string 

func main() {

	// atribuição do valor "João" à variável name
	name = "João" 

	// impressão do texto "Olá mundo! Meu nome é João." no console
	fmt.Println("Olá mundo! Meu nome é", name, "!")
}
```


### Múltiplas Declarações a nível de Package

É possível declarar várias variáveis a nível de package de uma só vez, usando a seguinte sintaxe:


```golang
var (
	
	name string
	n1 int
	n2 int
)
```

Vale observar que não é possível declarar variáveis com essa sintaxe no escopo de uma função.


### Atribuições a nível de função

Para atribuir um valor a uma variável a nível de função, use a seguinte sintaxe:

```golang
message := "Olá Mundo"
fmt.Println(message)
```

### Múltiplas Atribuições

É possível atribuir valores a várias variáveis de uma só vez, usando a seguinte sintaxe:


```golang
var a, b, c = true, 2.3, "Olá"
fmt.Println(a, b, c)
```

### Troca de valores entre variáveis

Para trocar os valores de duas variáveis, use a seguinte sintaxe:


```golang
// declaração das variáveis x e y
var x, y = 10, 20
fmt.Println("Antes da troca:", x, y)
// troca os valores das variáveis x e y
x, y = y, x
fmt.Println("Depois da troca:", x, y)
```

## Condicionais

As estruturas condicionais são usadas para controlar o fluxo de execução de um programa Go. No Go, existem dois tipos de estruturas condicionais: `if` e `switch`.


### IF

O condicional if é usado para verificar se uma condição é verdadeira ou falsa. Em go, o if não utiliza parentteses para delimitar a condição.

 A sintaxe do if é a seguinte:

```golang
package main

import "fmt"

func main() {

	a, b := 10, 15

	// Comparando dois números inteiros
	if a > b {
		println("a é maior que b")
	} else if a < b {
		println("a é menor que b")
	} else {
		println("a é igual a b")
	}
}
```


### Switch

O condicional switch em Go permite avaliar várias condições e executar o bloco de código correspondente. Diferentemente de algumas linguagens, Go não requer break após cada caso. O exemplo abaixo ilustra isso:

```golang
package main

import "fmt"

func main() {

	numero := 2

	switch numero {

	case 1:
		fmt.Println("O número é 1")

	case 2:
		fmt.Println("O número é 2")

	default:
		fmt.Println("Número não reconhecido")
	}
}
```

Neste exemplo simples de switch em Go:

numero é avaliado em cada caso:

- Se numero for 1, imprime "O número é 1".
- Se numero for 2, imprime "O número é 2".
- Se numero não corresponder a nenhum caso, o bloco default é executado, imprimindo "Número não reconhecido".

Não é necessário utilizar break após cada caso, pois o comportamento padrão do switch em Go é sair do switch após executar o caso correspondente.


### Switch com condicionais

Neste exemplo o código abaixo ilustra como usar um switch em vez de múltiplos if e else if:

```golang
package main

import "fmt"

func main() {
	a, b := 10, 15

	// Switch para comparar dois números inteiros
	switch {
	case a > b:
		fmt.Println("a é maior que b")
	case a < b:
		fmt.Println("a é menor que b")
	default:
		fmt.Println("a é igual a b")
	}
}
```


## Loops

Loops são usados para repetir um bloco de código um determinado número de vezes. No Go, existe loop `for`,porém diferente de outras linguagens, o Go não possui um loop `while`.

Abaixo, segue a forma mais simples de se usar um loop `for` em Go:

```golang
package main

func main() {

	for i := 0; i < 10; i++ {
		println(i)
	}
}
```

O código abaixo demonstra três maneiras de iterar um array de strings:


```golang
package main

func main() {

	names := []string{"Carlos", "Kelli", "Carol"}

	for i := 0; i < len(names); i++ {
		println(names[i])
	}

	for _, name := range names {
		println(name)
	}

	for i, name := range names {
		println(i, name)
	}
}
```

- O primeiro método **(for i := 0; i < len(names); i++)** é tradicional e útil quando você precisa acessar os índices do array.
- O segundo método **(for _, name := range names)** é conveniente quando você só precisa dos valores do array e não dos índices.
- O terceiro método **(for i, name := range names)** é ótimo para quando você precisa tanto dos índices quanto dos valores do array.

### Loop com Condição Simplificada: Controle manual do índice

No exemplo abaixo, usamos um loop for com uma condição simplificada. A variável i é declarada fora do loop e verificamos i < len(names) dentro do for. Em cada iteração, o elemento names[i] é impresso e i é incrementado. Isso oferece uma abordagem mais sucinta e fácil de entender:

```golang
package main

func main() {

	names := []string{"Carlos", "Kelli", "Carol"}

	var i int

	for i < len(names) {
		println(names[i])
		i++
	}
}
```

Este loop é útil quando queremos iterar sobre um array, mantendo o controle manual do índice. É uma alternativa clara e direta ao for tradicional com inicialização, condição e incremento.


### Loop for{} infinito

O laço `for {}` em Go é uma estrutura de loop infinito que é comumente usada em contextos de workers. Este tipo de loop é frequentemente utilizado para criar workers (trabalhadores) em Go, onde cada iteração do loop representa uma goroutine separada que executa uma determinada tarefa.

Por exemplo, em um sistema de processamento de filas, você pode ter um loop for {} que continua rodando para sempre. Dentro desse loop, você pode usar um canal para receber tarefas a serem executadas. Cada vez que uma tarefa é recebida do canal, uma nova goroutine é criada para processar essa tarefa.

A estrutura básica seria algo assim:

```golang
package main

import (
	"fmt"
	"time"
)

func worker() {
	i := 1

	// for infinito
	for {
		// Imprime a contagem atual
		fmt.Println("Contagem infinita:", i)
		i++

		// Simula o processamento contínuo
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("Iniciando worker...")

	// Inicia o worker em uma goroutine
	go worker()
	fmt.Println("Worker iniciado.")

	// Espera um pouco antes de encerrar o programa (para observar a contagem)
	time.Sleep(time.Second * 10)
	fmt.Println("Programa encerrado.")
}
```

## Tipos Complexos

Os tipos complexos em Go são estruturas de dados que podem armazenar múltiplos valores de tipos diferentes. São amplamente utilizados para modelar e organizar dados de maneira eficiente e flexível, oferecendo recursos poderosos para lidar com diferentes cenários de programação. 

Alguns exemplos são:

1. **Array:** Uma coleção fixa de elementos do mesmo tipo, com tamanho definido na declaração. 
Exemplo:

```golang
numeros := [5]int{1, 2, 3, 4, 5}
```

2. **Slice:** Uma "fatia" dinâmica de um array, permitindo um tamanho flexível. 
Exemplo:

```golang
numeros := []int{1, 2, 3, 4, 5}
```

3. **Map:** Uma coleção de pares chave-valor, onde cada chave é única. 
Exemplo:

```golang
pessoas := map[string]int{"João": 30, "Maria": 25, "Pedro": 28}
```

4. **Struct:** Uma estrutura de dados que agrupa campos de diferentes tipos. 
Exemplo:

```golang
type Pessoa struct {
    Nome string
    Idade int
}

pessoa := Pessoa{"Ana", 35}
```

### Array 

Em Go, um `array` é uma estrutura de dados estática que representa uma coleção fixa de elementos do mesmo tipo. Ele é definido com um tamanho específico durante a declaração e não pode ser alterado após a criação.

```golang
package main

import "fmt"

func main() {

	// Cria um array de tamanho fixo de strings
	var nomes [3]string

	// Atribui valores aos elementos do array
	nomes[0] = "João"
	nomes[1] = "Maria"
	nomes[2] = "Pedro"

	// Imprime o array
	fmt.Println(nomes)

	// Acessando elementos do array
	fmt.Println(nomes[0]) // Imprime "João"
}
```

Neste exemplo, o array nomes é declarado com tamanho fixo de 3 elementos do tipo string. Os valores são atribuídos aos elementos individualmente. Como o array tem tamanho fixo, não podemos adicionar ou remover elementos após a sua criação. O acesso aos elementos é feito utilizando índices, como nomes[0] para acessar o primeiro elemento.

**Resumo:**

- Um array tem um tamanho fixo e é declarado com [Tamanho]Tipo.
- Os elementos de um array são acessados por índices.
- É estático e não pode ser alterado após a criação.
- Útil quando o tamanho é conhecido e não muda.

### Slice

Em Go, um `slice` slice é uma estrutura de dados dinâmica que representa uma fatia de um array. Ele possui um tamanho dinâmico e uma capacidade (ou tamanho do array subjacente), permitindo que cresça conforme necessário e pode ser considerado como tendo um tamanho "infinito" em termos de sua capacidade de expansão, pois o Go gerencia automaticamente a alocação de memória conforme mais elementos são adicionados. Isso permite que o slice cresça conforme necessário, sem a necessidade de especificar previamente seu tamanho máximo.


```golang
package main

import (
	"fmt"
)

func main() {

	// Cria um slice de strings
	names := []string{"Tiago", "Daniel", "João"}
	fmt.Println(names)

    // Adiciona um novo elemento ao slice
	names = append(names, "Pedro")
	fmt.Println(names)

    
	// Adiciona um novo elemento ao slice, 
	names = append(names, "João")
	fmt.Println(names, len(names), cap(names))
}
```

Conforme o exemplo de código acima:

- `names := []string{"Tiago", "Daniel", "João"}` cria um slice de strings.
- `append(names, "Pedro")` adiciona "Pedro" ao slice, automaticamente expandindo conforme necessário.
- `append(names, "João")` adiciona "João" ao slice, mostrando sua capacidade dinâmica.
- `len(names)` retorna o comprimento atual do slice (5 elementos).
- `cap(names)` retorna a capacidade atual (8 elementos, expandindo conforme necessário).


### Array vs Slice

**Array:**

- Um array em Go é uma coleção fixa de elementos do mesmo tipo.
- Seu tamanho é definido na declaração e não pode ser alterado.
- A declaração de um array seria assim:

```golang
var nomes [3]string // Array de strings com tamanho 3
nomes[0] = "João"
nomes[1] = "Maria"
nomes[2] = "Pedro"
```

**Slice:**

- Um slice em Go é uma "fatia" dinâmica de um array, permitindo um tamanho flexível.
- Não precisa de um tamanho definido na declaração.
- A declaração de um slice seria assim:

```golang
nomes := []string{"João", "Maria", "Pedro"} // Slice de strings
```

### Map

Em Go, um `map` é uma estrutura de dados que mapeia chaves únicas para valores associados. Ele é semelhante a um dicionário em outras linguagens. 

```golang
package main

import (
	"fmt"
)

func main() {

	// Create a map of names and ages
	people := make(map[string]int32)
	people["João"] = 45
	people["Maria"] = 42
	people["Antônio"] = 39

	// Print the map
	fmt.Println("People:", people)
	// output
	// People: map[João:45 Maria:42 Antônio:39]

	// Get the age of Maria
	name, ok := people["Maria"]
	fmt.Println(name, ok)
	// output
	// Maria 42 true

	// Try to get the age of Pedro
	name, ok = people["Pedro"]
	fmt.Println(name, ok)
	// output
	// 0 false
}
```

- `people := make(map[string]int32)` cria um map onde as chaves são strings e os valores são inteiros de 32 bits.
- `people["João"] = 45, people["Maria"] = 42, people["Antônio"] = 39` adicionam valores ao map associados às chaves correspondentes.
- `fmt.Println("People:", people)` imprime o map completo.
  - **Saída:** `People: map[João:45 Maria:42 Antônio:39]`
- `name, ok := people["Maria"]` obtém o valor associado à chave "Maria" e verifica se a chave existe.
  - **Saída:** `Maria 42 true`
- `name, ok = people["Pedro"]` tenta obter o valor associado à chave "Pedro", mas como essa chave não existe, o valor zero para o tipo (0 para int32) é retornado e ok é false.
  - **Saída**: `0 false`

 Em resumo, um map em Go é uma coleção de pares chave-valor onde as chaves são únicas e os valores podem ser recuperados rapidamente com base na chave. O uso do map permite uma busca eficiente e fácil acesso aos valores associados às chaves.


### Struct 

Em Go, uma `struct` é uma coleção de campos que agrupa dados de diferentes tipos relacionados. A técnica para atributos públicos e privados em uma struct é determinada pela letra inicial do nome do campo:

```golang
package main

import "fmt"

type Client struct {
	Name   string
	age    int
	Email  string
	Phone  string
	Status bool
}

func main() {

	// Cria um novo cliente com os valores especificados
	client := Client{
		Name:   "John Doe",
		age:    30,
		Email:  "john.doe@example.com",
		Phone:  "123-456-7890",
		Status: true, // O Go exige que sempre termine com vírgula
	}

	fmt.Println(client) // Imprime o cliente no console

	fmt.Println(client.Name, client.Email) // Imprime o nome e o e-mail do cliente no console
}
```

- `type Client struct { ... }` define a estrutura Client com campos para informações de um cliente.
- `client := Client{ ... }` cria um novo cliente com valores específicos para cada campo.
- `fmt.Println(client)` imprime todos os campos do cliente no console.
  - **Saída:** `{John Doe 30 john.doe@example.com 123-456-7890 true}`
- `fmt.Println(client.Name, client.Email)` imprime o nome e o e-mail do cliente.
  - **Saída:** `John Doe john.doe@example.com`

Neste exemplo, client é uma instância da estrutura Client, que funciona como um "array" de campos com nome e tipo definidos. Cada campo pode ser acessado usando o ponto (.) seguido pelo nome do campo (client.Name, client.Email, etc.).


#### Auto Relacionamento - Ponteiro

Trata-se da técnica de autorelacionamento por ponteiro, onde uma estrutura (struct) pode conter um campo que é um ponteiro para outra instância da mesma estrutura. 

Isso permite criar relações hierárquicas dentro da mesma estrutura de dados, como um "pai" que pode ter um "filho" que também é do mesmo tipo. O campo Father na estrutura Client é um exemplo de autorelacionamento usando ponteiro.

```golang
package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Email  string
	Phone  string
	Status bool
	Father *Client // Autorelacionamento: Ponteiro para outro Client
}

func main() {

	// Criando instâncias de Client
	child := Client{
		Name:   "João",
		Age:    30,
		Email:  "joao@example.com",
		Phone:  "123-456-7890",
		Status: true,
	}

	father := Client{
		Name:   "Pedro",
		Age:    25,
		Email:  "pedro@example.com",
		Phone:  "987-654-3210",
		Status: true,
	}

	// Estabelecendo a relação "pai-filho"
	child.Father = &father

	// Imprimindo informações
	fmt.Println("Cliente 1:")
	fmt.Println("Nome:", child.Name)
	fmt.Println("Idade:", child.Age)
	fmt.Println("E-mail:", child.Email)
	fmt.Println("Telefone:", child.Phone)
	fmt.Println("Status:", child.Status)
	if child.Father != nil {
		fmt.Println("Pai:", child.Father.Name)
	} else {
		fmt.Println("Pai: Nenhum")
	}

	fmt.Println("\nCliente 2:")
	fmt.Println("Nome:", father.Name)
	fmt.Println("Idade:", father.Age)
	fmt.Println("E-mail:", father.Email)
	fmt.Println("Telefone:", father.Phone)
	fmt.Println("Status:", father.Status)
}
```

Neste exemplo, criamos duas instâncias de `Client` chamadas `child` e `father`. O `child` é "João" e o `father` é "Pedro". Em seguida, estabelecemos a relação "pai-filho" onde child aponta para father como seu "pai". 

Ao imprimir as informações, podemos ver que `child` tem um "pai" definido como `father`, demonstrando o autorelacionamento usando um ponteiro para outra instância da mesma estrutura.


#### Visibilidade

- Campos começando com letra maiúscula são públicos e podem ser acessados fora do pacote.
- Campos começando com letra minúscula são privados e só podem ser acessados dentro do mesmo pacote.

Conforme o exemplo de código acima, `Name` é um campo público, permitindo acesso fora da struct, enquanto `age` é privado e só pode ser acessado dentro do mesmo pacote. É uma técnica comum em Go para controlar a visibilidade e encapsulamento dos dados em uma struct.


## Funções

Funções em Go são blocos de código independentes que podem ser chamados em qualquer parte do programa, e são definidas fora de qualquer estrutura ou tipo e podem ser chamadas com argumentos.

Para declarar uma função, use a seguinte sintaxe:


```golang
package main

import (
	"fmt"
)

func hello(nome string) {
	fmt.Println("Olá", nome, "!")
}

func sum(a, b int) int {
	return a + b
}

func main() {

	hello("João")

	fmt.Println("Sua idade é:", sum(10, 20), "anos")
}
```


### Funções com múltiplos retornos

Em Go é comum usar funções com mais de um retorno, especialmente para lidar com possíveis erros. Por exemplo:

```golang
package main

import (
	"errors"
	"fmt"
)

// Função divide recebe dois inteiros e retorna um inteiro e um erro
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}
	return a / b, nil
}

// A função main é onde a execução do programa começa.
func main() {
	// Chamando a função divide com 10 como dividendo e 2 como divisor
	resultado, err := divide(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado da divisão:", resultado)
	}

	// Chamando a função divide com 10 como dividendo e 0 como divisor
	resultado, err = divide(10, 0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado da divisão:", resultado)
	}
}
```

Neste exemplo, a função **divide** tem dois retornos: resultado do tipo int e erro do tipo error. Se o divisor for zero, ela retorna um erro. No main, verificamos se há um erro após chamar a função divide e imprimimos a mensagem de erro. Caso contrário, imprimimos o resultado da divisão.

## Métodos

Métodos em Go são funções associadas a um tipo específico, chamado receptor (receiver). São definidos dentro do bloco de declaração de um tipo e têm acesso aos campos desse tipo. Métodos são chamados em uma instância (ou ponteiro) do tipo ao qual estão associados.

```golang
package main

import "fmt"

type Client struct {
	Name string
	Age  int
}

// Método de Pessoa para exibir os detalhes
func (c Client) ShowDetails() {
	fmt.Println("Nome:", c.Name)
	fmt.Println("Idade:", c.Age)
}

func main() {
	client := Client{Name: "Ana", Age: 30}
	client.ShowDetails()
}
```

No exemplo acima, `ShowDetails` é um método associado ao tipo cliente. Ele pode ser chamado em uma instância de `Client` para exibir os detalhes dessa pessoa. Em resumo, os métodos são como funções específicas de um tipo, enquanto as funções em Go podem ser chamadas independentemente de qualquer tipo.


### Métodos com Ponteiro

Em Go, métodos com ponteiro permitem modificar diretamente a estrutura original em que são chamados. Isso é particularmente útil ao lidar com estruturas de dados complexas ou quando queremos preservar as mudanças feitas nos métodos.

Vamos considerar a estrutura `Category` e os métodos `HasParent` e `SetFather`:

```golang
package main

import "fmt"

type Category struct {
	Name   string
	Parent *Category
}

func (c Category) HasParent() bool {
	return c.Parent != nil
}

func (c *Category) SetFather(father *Category) {
	c.Parent = father
}

func main() {
	// Criando uma categoria "Notebooks" sem pai
	notebooks := Category{Name: "Notebooks", Parent: nil}

	// Verificando se "Notebooks" tem um pai (deve retornar false)
	fmt.Println("Notebooks tem pai?", notebooks.HasParent())

	// Criando uma categoria "Tecnologia" como pai de "Notebooks"
	tecnologia := Category{Name: "Tecnologia", Parent: nil}
	notebooks.SetFather(&tecnologia)

	// Verificando se "Notebooks" tem um pai agora (deve retornar true)
	fmt.Println("Notebooks tem pai?", notebooks.HasParent())
	fmt.Println("Pai de Notebooks:", notebooks.Parent.Name) // Imprime "Tecnologia"
}
```

Neste exemplo, `HasParent` verifica se uma categoria tem um pai, enquanto `SetFather` atribui um pai a uma categoria. Ambos os métodos são definidos com um receptor de ponteiro (*Category), permitindo que modifiquem diretamente a estrutura Category original em que são chamados.


## Herança (Embeding)

Em Go, não existe o conceito de herança da forma tradicional presente em algumas linguagens de programação orientada a objetos. Em vez disso, Go utiliza composição e interfaces para alcançar reutilização de código e polimorfismo.


### Composição em Go

- Em Go, é comum usar a composição para reutilizar funcionalidades de outros tipos.
- A composição permite que um tipo contenha outro tipo como um campo.
- Não há herança de membros de estrutura, mas sim incorporação (embedding) de tipos.


```golang
package main

import "fmt"

// Definindo um tipo Veiculo básico
type Veiculo struct {
	Modelo string
	Ano    int
}

// Método para parar um veículo
func (v Veiculo) Parar() {
	fmt.Printf("O veículo %s está parado.\n", v.Modelo)
}

// Definindo um tipo Carro que incorpora Veiculo
type Carro struct {
	Veiculo // Incorporação do tipo Veiculo
	Motor   string
}

// Método para o carro andar
func (c Carro) Andar() {
	fmt.Printf("O carro %s está andando com seu motor %s...\n", c.Modelo, c.Motor)
}

// Definindo um tipo Barco que incorpora Veiculo
type Barco struct {
	Veiculo // Incorporação do tipo Veiculo
	Helice  string
}

// Método para o barco navegar
func (b Barco) Navegar() {
	fmt.Printf("O barco %s está navegando com sua hélice %s...\n", b.Modelo, b.Helice)
}

func main() {
	
	// Criando um Carro
	carro := Carro{
		Veiculo: Veiculo{Modelo: "Fusca", Ano: 1972},
		Motor:   "V8",
	}

	// Criando um Barco
	barco := Barco{
		Veiculo: Veiculo{Modelo: "Lancha", Ano: 2020},
		Helice:  "Marinha 3000",
	}

	// Chamando o método Andar para o Carro
	carro.Andar()

	// Chamando o método Navegar para o Barco
	barco.Navegar()

	// Chamando o método Parar para ambos os veículos
	carro.Parar()
	barco.Parar()
}
```

Neste exemplo, criamos os tipos Veiculo, Carro, e Barco, onde Carro e Barco herdam características do tipo Veiculo. Além disso, cada tipo possui atributos específicos que representam suas características únicas.

- `Veiculo`: É um tipo básico que representa as características comuns de um veículo, como o modelo e o ano.

- `Carro`: É um tipo que incorpora Veiculo e também possui um atributo específico, Motor, que representa o motor do carro.

- `Barco`: Também incorpora Veiculo e possui um atributo específico, Helice, que representa a hélice do barco.

- Os métodos `Andar` e `Navegar` são específicos para `Carro` e `Barco`, respectivamente. Eles demonstram a ação de andar para um carro e navegar para um barco, utilizando seus atributos específicos.

- O método `Parar` é comum a todos os tipos, pois um veículo de qualquer tipo pode ser parado.

No `main`, criamos uma instância de `Carro` e `Barco` com seus respectivos atributos específicos (`Motor` e `Helice`), e demonstramos o uso dos métodos `Andar`, `Navegar`, e `Parar` para cada veículo. Esta abordagem de herança e atributos específicos é bastante útil para representar objetos do mundo real de forma clara e eficiente em Go.


## Interface

Em Go, interfaces são conjuntos de métodos que definem um comportamento. Assim como nas demais linguagens, uma interface define um contrato que especifica quais métodos uma struct deve implementar para satisfazer a interface. Isso permite que diferentes tipos possam ser tratados de maneira uniforme se implementarem os métodos da interface.

Um exemplo simples de uma interface em Go:


```golang
package main

import "fmt"

// Definindo a interface Animal
type Animal interface {
	Som() string
	Andar() string
}

// Definindo uma struct Dog que implementa a interface Animal
type Dog struct {
	Nome string
}

// Implementação do método Som para Dog
func (d Dog) Som() string {
	return "Au Au"
}

// Implementação do método Andar para Dog
func (d Dog) Andar() string {
	return fmt.Sprintf("%s está caminhando...", d.Nome)
}

func main() {
	// Criando uma instância de Dog
	cachorro := Dog{Nome: "Rex"}

	// Usando métodos da interface Animal
	fmt.Println("Som do", cachorro.Nome+":", cachorro.Som())
	fmt.Println("Movimento do", cachorro.Nome+":", cachorro.Andar())
}
```

Neste exemplo:

- `Animal` é uma interface que define dois métodos: `Som` e `Andar`.
- `Dog` é uma struct que implementa a interface `Animal` ao definir os métodos `Som` e `Andar`.
- No `main`, criamos um Dog chamado `cachorro` e chamamos seus métodos `Som` e `Andar`, que são métodos da interface `Animal`. 

Este exemplo demonstra como uma interface em Go permite que diferentes tipos (como `Dog`, `Cat`, `Bird`, etc.) podem ser tratados de forma genérica se implementarem os métodos da interface. Isso proporciona flexibilidade e reutilização de código.


## Asserções de Interfaces

As asserções de interfaces em Go permitem verificar e converter um valor de interface para um tipo subjacente. É uma forma de verificar se uma interface contém um tipo específico e, se sim, obter o valor subjacente desse tipo. Isso é útil para lidar com tipos dinâmicos de forma segura.

```golang
package main

import "fmt"

// Definindo a interface Animal
type Animal interface {
	Som() string
	Andar() string
}

// Definindo uma struct Dog que implementa a interface Animal
type Dog struct {
	Nome string
}

// Definindo uma struct Cat que implementa a interface Animal
type Cat struct {
	Nome string
}

// Implementação do método Som para Dog
func (d Dog) Som() string {
	return "Au Au"
}

// Implementação do método Andar para Dog
func (d Dog) Andar() string {
	return fmt.Sprintf("%s está caminhando...", d.Nome)
}

// Implementação do método Som para Cat
func (c Cat) Som() string {
	return "Miau Miau"
}

// Implementação do método Andar para Cat
func (c Cat) Andar() string {
	return fmt.Sprintf("%s está caminhando...", c.Nome)
}

// Verifica o tipo subjacente da interface Animal
func show(a Animal) {

	// Condicional switch com type para verificar o tipo subjacente da interface Animal passada como argumento
	switch a.(type) {
	case Dog:
		fmt.Println("Som do", a.(Dog).Nome+":", a.(Dog).Som())
		fmt.Println("Movimento do", a.(Dog).Nome+":", a.(Dog).Andar())
	case Cat:
		fmt.Println("Som do", a.(Cat).Nome+":", a.(Cat).Som())
		fmt.Println("Movimento do", a.(Cat).Nome+":", a.(Cat).Andar())
	default:
		fmt.Println("Animal desconhecido.")
	}
	fmt.Println()
}

func main() {

	// Criando uma instância de Dog
	cachorro := Dog{Nome: "Rex"}
	show(cachorro)

	// Criando uma instÂncia de Cat
	gato := Cat{Nome: "Bichano"}
	show(gato)
}
```

Neste exemplo:

- No método `show`, usamos um switch com `type` para verificar o tipo subjacente da interface `Animal` passada como argumento.
- Para cada caso (`Dog` ou `Cat`), realizamos a asserção de tipo `(a.(Dog))` ou `(a.(Cat))` para converter a interface para o tipo desejado.
- Isso nos permite acessar os métodos específicos de `Dog` ou `Cat` (como `Som` e `Andar`) e imprimir suas características.
- Este método é útil quando temos uma interface e queremos executar comportamentos específicos com base no tipo subjacente da interface.


## Concorrência vs Paralelismo

Ao trabalhar com Go, é importante compreender a diferença entre concorrência e paralelismo. Embora os termos possam parecer semelhantes, eles têm significados distintos no contexto da linguagem.

### Concorrência

Em Go, a concorrência refere-se à capacidade de um programa executar várias tarefas ao mesmo tempo. Isso é alcançado usando **[goroutines](https://golang.org/doc/articles/concurrency.html)**, que são threads leves gerenciadas pelo Go runtime. Uma goroutine permite que você inicie uma função de forma independente e ela será executada de forma assíncrona em relação ao restante do programa. Isso é útil para processamento concorrente, como manipulação de requisições em servidores web ou execução de várias tarefas ao mesmo tempo.

### Paralelismo

Por outro lado, o paralelismo refere-se à capacidade de executar várias tarefas simultaneamente em CPUs múltiplos ou núcleos de CPU. Isso significa que várias partes do programa estão sendo executadas verdadeiramente ao mesmo tempo, proporcionando uma maneira de melhorar o desempenho ao lidar com tarefas intensivas em CPU.

### Como o Go lida com isso

O Go é conhecido por sua abordagem de concorrência em vez de paralelismo. Ele permite que você crie facilmente goroutines para executar tarefas de forma concorrente, aproveitando eficientemente o poder de processamento disponível. 

Em resumo, o Go é projetado para ser concorrente, aproveitando as goroutines para realizar várias tarefas ao mesmo tempo em um único núcleo ou em vários núcleos.


## Go Routines

As goroutines são uma forma de executar funções de forma assíncrona e concorrente, semelhante a threads, mas com um custo de criação muito menor. 

- São "threads leves" gerenciadas pelo próprio runtime do Go.
- São unidades de execução leves que permitem que um programa execute funções de forma assíncrona e concorrente.

Podemos entender as goroutines como processos independentes que executam tarefas em segundo plano, permitindo que o programa continue executando outras operações.

### Principais Características:

- **Leveza:** Uma goroutine é leve em comparação com uma thread tradicional do sistema operacional. Isso significa que você pode criar milhares (ou até mesmo milhões) de goroutines em um programa Go sem consumir muitos recursos.
- **Concorrência Simples:** O Go facilita a criação e execução de goroutines. Basta adicionar a palavra-chave go antes de uma função para executá-la como uma goroutine.
- **Comunicação:** As goroutines podem se comunicar entre si usando canais (`channels`). Isso permite a sincronização e a troca de dados entre goroutines de forma segura.


```golang
package main

import (
	"fmt"
	"time"
)

func tarefa1() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Executando tarefa 1...", i)
		time.Sleep(time.Second) // Simula uma tarefa que leva 1 segundo
	}
}

func tarefa2() {
	for i := 1; i <= 3; i++ {
		fmt.Println("Executando tarefa 2...", i)
		time.Sleep(800 * time.Millisecond) // Simula uma tarefa que leva 800 milissegundos
	}
}

func main() {
	fmt.Println("Iniciando programa principal.")

	// Executa a função tarefa1 como uma goroutine
	go tarefa1()

	// Executa a função tarefa2 como uma goroutine
	go tarefa2()

	fmt.Println("Continuando a execução do programa principal.")

	// Aguarda um pouco para não terminar imediatamente
	time.Sleep(4 * time.Second)

	fmt.Println("Programa principal finalizado.")
}
```

Neste exemplo, `tarefa1()` e `tarefa2()` são executadas como goroutines simultaneamente. Enquanto uma está sendo executada, a outra também está em andamento, permitindo a execução concorrente das duas tarefas.

Vale lembrar que a função `main` também é uma goroutine.


## Channels

Channels em Go são uma poderosa ferramenta para a comunicação e sincronização entre goroutines. Eles fornecem uma maneira segura de transmitir dados de uma goroutine para outra. 

- São canais de comunicação que conectam goroutines.
- Permitem que goroutines enviem e recebam valores de forma sincronizada.
- Seguem um modelo de comunicação baseado em "enviar" e "receber".


### Características Principais:

- **Comunicação Segura:** Channels garantem que a comunicação entre goroutines seja segura e sem race conditions (concorrência de dados).
- **Blqueio:** As operações de leitura (`<-`) e escrita (`<-`) em um channel são bloqueantes. Isso significa que uma goroutine será bloqueada até que a outra esteja pronta para receber ou enviar dados.
- **Sincronização:** Channels podem ser usados para sincronizar a execução de goroutines. Por exemplo, uma goroutine pode aguardar até receber um sinal de outra goroutine através de um channel.


### Sintaxe Básica:


#### Criar um channel

```golang
ch := make(chan TipoDoDado)
```

#### Enviar valor para um channel

```golang
ch <- valor
```

#### Receber valor para um channel

```golang
valor := <-ch
```

#### Exemplo de uso

```golang
package main

import "fmt"

func enviarDados(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Envia valores para o channel
		fmt.Println("Valor enviado:", i)
	}
	close(ch) // Fecha o channel após enviar todos os valores
}

func receberDados(ch <-chan int) {
	for {
		valor, ok := <-ch // Recebe valores do channel
		if !ok {
			fmt.Println("Channel fechado.")
			return
		}
		fmt.Println("Valor recebido:", valor)
	}
}

func main() {
	// Criar um channel
	dados := make(chan int)

	// Executar a função para enviar dados em uma goroutine
	go enviarDados(dados)

	// Executar a função para receber dados em outra goroutine
	go receberDados(dados)

	// Aguardar um pouco para não terminar imediatamente
	fmt.Println("Aguardando...")
	fmt.Scanln()
}
```

Neste exemplo:

- `enviarDados` envia valores de 1 a 5 para o channel dados e fecha o channel após enviar todos os valores.
- `receberDados` recebe os valores do channel dados e os imprime.
- As duas goroutines (`enviarDados` e `receberDados`) executam simultaneamente, sincronizadas pelo channel.

Isso demonstra como os channels podem ser usados para transmitir dados de forma segura e sincronizada entre goroutines em Go.

### Tipos de Channels

#### 1. Channel de Leitura (`<-chan TipoDoDado`):

- Define um canal somente para leitura.
- Apenas permite operações de recebimento (`valor := <-ch`), não é possível enviar valores para este canal.
- Útil quando queremos garantir que nenhuma goroutine acidentalmente envie valores para o canal.

Exemplo:

```golang
ch := make(<-chan int) // Canal de leitura
```

#### 2. Channel de Escrita (`chan<- TipoDoDado`):

- Define um canal somente para escrita.
- Apenas permite operações de envio (`ch <- valor`), não é possível receber valores deste canal.
- Útil quando queremos garantir que nenhuma goroutine acidentalmente leia valores do canal.

Exemplo:

```golang
ch := make(chan<- int) // Canal de escrita
```

#### 3. Channel de Leitura e Escrita (`chan TipoDoDado`):

- Define um canal para leitura e escrita.
- Permite tanto operações de envio (`ch <- valor`) quanto de recebimento (`valor := <-ch`).
- É o tipo de canal mais comum e versátil.

Exemplo:

```golang
ch := make(chan int) // Canal de leitura e escrita
```

Estes tipos de channels são úteis para diferentes cenários. Canais de leitura (`<-chan`) e escrita (`chan<-`) são usados quando queremos restringir o uso do canal para apenas uma operação específica, garantindo mais segurança e clareza no código. Já canais de leitura e escrita (`chan`) são os mais utilizados, pois oferecem a capacidade de envio e recebimento de valores, sendo essenciais para a comunicação entre goroutines em Go.


## Gerando o arquivo binário

Para compilar um arquivo Go e gerar o arquivo binário executável, você pode usar o comando `go build` seguido do nome do arquivo. Por exemplo, se o seu arquivo Go se chama `main.go`, você pode usar o seguinte comando no terminal:

```bash
go build main.go
```

Este comando compilará o arquivo `main.go` e criará um arquivo binário executável chamado `main` (ou `main.exe` no Windows) no mesmo diretório onde o arquivo Go está localizado.


### Executando o arquivo binário

Para executar o arquivo binário gerado após a compilação, você pode simplesmente digitar o nome do arquivo no terminal (ou prompt de comando no Windows) precedido de `./` no Linux e macOS, ou apenas o nome no Windows. Por exemplo:

No Linux e macOS:

```bash
./main
```

No Linux e macOS:

```cmd
main.exe
```

Isso executará o arquivo binário `main`, que é o programa que você compilou a partir do arquivo Go. Você verá a saída do programa no terminal. Certifique-se de estar no mesmo diretório em que o arquivo binário foi criado ou especificar o caminho completo para o arquivo se estiver em um diretório diferente.


## Conclusão

Durante este estudo, exploramos os fundamentos da linguagem de programação Go (Golang), desde os conceitos básicos como variáveis, funções, e estruturas de controle, até tópicos mais avançados como concorrência, canais, e interfaces. Aqui estão alguns pontos-chave:

- **Simplicidade e Eficiência:** Go foi projetada para ser uma linguagem simples, clara e eficiente, com sintaxe concisa e poderosa.
- **Concorrência e Goroutines:** A concorrência em Go é impulsionada pelas goroutines, que são leves e podem ser executadas em paralelo. Isso permite que escrevamos programas concorrentes de forma simples e eficiente.
- **Canais (Channels):** Os canais são a principal forma de comunicação entre goroutines em Go. Eles permitem a sincronização e a troca de dados de forma segura e eficiente.
- **Tipos de Channels:** Existem canais de leitura (`<-chan`), escrita (`chan<-`), e leitura e escrita (`chan`) para atender às diferentes necessidades de comunicação entre goroutines.
- **Interfaces:** As interfaces em Go permitem a abstração de comportamentos comuns, permitindo a implementação de tipos diferentes com métodos similares. Isso promove o polimorfismo e a reutilização de código.
- **Estruturas e Métodos:** Utilizamos structs para representar dados estruturados e métodos para associar comportamentos específicos a esses tipos de dados.
- **Compilação e Execução:** Para compilar um programa Go, usamos o comando `go build`. Para executar o programa gerado, basta chamar o arquivo binário resultante.

Em resumo, Go é uma linguagem poderosa, projetada para simplicidade, eficiência e concorrência. Com sua sintaxe limpa e recursos avançados, é uma excelente escolha para desenvolvimento de software moderno e escalável.



