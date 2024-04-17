# Gin

O Gin é um framework web para Go (Golang), é usado para criar APIs web de forma rápida e simples em Go. O Gin fornece uma série de funcionalidades úteis e um conjunto de middlewares que ajudam os desenvolvedores a lidar com tarefas comuns ao criar aplicativos web, como roteamento, renderização de JSON, logging, autenticação, entre outros.

Com o Gin, os desenvolvedores podem criar endpoints de API de forma bastante eficiente, lidar com solicitações HTTP, validar dados de entrada e muito mais. É uma ferramenta popular para desenvolvedores que desejam construir aplicativos web em Go de maneira eficaz e com uma curva de aprendizado relativamente baixa.


## Inicializando um Módulo

Antes de tudo, é necessário inicializar um módulo para o  projeto. Podemos fazer isso executando o seguinte comando na pasta onde está o seu código:

```bash
go mod init nome_do_seu_modulo
```

Este comando criará um arquivo chamado `go.mod` que registra o nome do seu módulo e suas dependências.


## Instalando o Gin

Antes de tudo, é necessário ter o `Gin` instalado usando o comando:

```bash
go get -u github.com/gin-gonic/gin
```

Este comando baixará e instalará o Gin e suas dependências.


## Liberando a Porta 8080

Caso a porta utilizada fique presa no processo, utilize o comando abaixo para liberar:

```bash
kill -9 $(lsof -t -i:8081)
```