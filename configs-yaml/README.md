# Configurações com YAML no Go: Leitura, Deserialização e Iteração

Se precisa lidar com configurações de forma dinâmica, aprender a utilizar arquivos YAML é uma excelente escolha. 

Eles são simples, legíveis e ideais para armazenar configurações da sua aplicação. Neste breve estudo, vamos explorar como ler arquivos YAML no Go e aplicar essas configurações em seu código aproveitando o pacote `gopkg.in/yaml.v3`.

## O que é YAML?

YAML (YAML Ain't Markup Language) é um formato de dados que se destaca por sua simplicidade e legibilidade. 

Ele é amplamente utilizado em arquivos de configuração, devido à sua estrutura intuitiva e fácil de manter.

```yaml
http_port:     "8080"
db_engineer:   "postgres"
db_connection: "user:password@/dbname"
table_name:    "tb_users"
```

Este arquivo define algumas configurações básicas de uma aplicação, como a porta HTTP e informações de conexão com o banco de dados.

## Passo a Passo para Usar YAML no Golang

Aqui está o código completo para ler as configurações a partir de um arquivo YAML no Golang:

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Struct para receber as configurações mapeadas do arquivo YAML
type Config struct {
	HTTPPort     string `yaml:"http_port"`
	DBEngineer   string `yaml:"db_engineer"`
	DBConnection string `yaml:"db_connection"`
	TableName    string `yaml:"table_name"`
}

func main() {
	// Leitura do arquivo YAML
	data, err := ioutil.ReadFile("configs/config.yaml")
	if err != nil {
		log.Fatalf("Erro ao ler arquivo YAML: %v", err)
	}

	// Desserialização do arquivo YAML via Unmarshal para a struct Config
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Erro ao fazer unmarshal do YAML: %v", err)
	}

	// Exibição dos valores
	fmt.Printf("Porta HTTP: %s\n", config.HTTPPort)
	fmt.Printf("Banco de Dados: %s\n", config.DBEngineer)
	fmt.Printf("Conexão: %s\n", config.DBConnection)
	fmt.Printf("Nome da tabela: %s\n", config.TableName)
}
```
### Explicação do Código

1. **Estrutura de Configuração:** Definimos uma struct Config para armazenar os dados lidos do arquivo YAML. Cada campo da struct tem uma tag associada ao nome da chave YAML correspondente.

2. **Leitura do Arquivo YAML:** O arquivo é lido usando a função `ioutil.ReadFile`. Se houver qualquer erro, ele será registrado e a execução será interrompida.

3. **Deserialização (Unmarshal):** O conteúdo do arquivo YAML é convertido para a estrutura Config usando a função `yaml.Unmarshal`.

4. **Exibição das Configurações:** Por fim, os valores lidos do YAML são exibidos no terminal, apenas como demonstração.


### Exemplo do Arquivo YAML que utilizaremos

Aqui está um exemplo prático do arquivo YAML usado neste código:

```yaml
http_port:     "8080"
db_engineer:   "postgres"
db_connection: "user:password@/dbname"
table_name:    "tb_users"
```

Esse arquivo simula as configurações para porta HTTP da aplicação, o tipo de banco de dados (neste caso, PostgreSQL), a string de conexão com o banco de dados, e o nome da tabela que será usada.

### Onde Posicionar o Arquivo YAML?

Para este exemplo, o arquivo `config.yaml` será armazenado em uma pasta chamada **configs** no diretório raiz do nosso projeto. A hierarquia ficaria assim:

```bash

/configs
    └── config.yaml
/main.go
```

### Explicando a Dependência **gopkg.in/yaml.v3**

Para integrar de arquivos YAML no Go, utilizamos o pacote `gopkg.in/yaml.v3`, que permite ler, escrever e manipular dados YAML de forma simples e eficiente.

**Instalação:** Execute o seguinte comando no terminal para adicionar a dependência ao projeto:

```bash
go get gopkg.in/yaml.v3
```

#### Alternativas ao gopkg.in/yaml.v3

Existem outras bibliotecas que podem ser úteis dependendo do que você precisa:

- `spf13/viper`: Esta é uma biblioteca mais completa, que além de suportar YAML, trabalha com outros formatos como JSON e TOML, além de variáveis de ambiente.

- `goccy/go-yaml`: sta biblioteca é uma excelente opção com foco em desempenho e recursos avançados de serialização e deserialização de YAML.

### Como Executar a Aplicação

Depois de configurar o código Go e o arquivo YAML, chegou a hora de executar a aplicação para ver os resultados.

Certifique-se de estar na raiz do seu projeto, onde o arquivo `main.go` está localizado.
    
Execute o seguinte comando no terminal:

```bash
go run main.go
```

Se tudo estiver corretamente configurado, os valores das configurações YAML serão impressos no terminal. O resultado deve ser algo parecido com isso:

```bash
Porta HTTP: 8080
Banco de Dados: postgres
Conexão: user:password@/dbname
Nome da tabela: tb_users
```

Esse comando compila e executa seu código, lendo as configurações definidas no arquivo `config.yaml` e exibindo-as no terminal.


## Conclusão

Neste estudo, foi apresentado como carregar e utilizar configurações armazenadas em arquivos YAML com Go. Usando o pacote gopkg.in/yaml.v3, conseguimos transformar um arquivo YAML em uma estrutura Go, pronta para ser utilizada. 

Isso é particularmente útil em aplicações que precisam de flexibilidade para ajustar configurações sem a necessidade de recompilar o código. Com esse conhecimento, você está pronto para integrar YAML nas suas aplicações Golang!