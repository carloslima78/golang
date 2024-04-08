# SQLite3

O `sqlite3` é um banco de dados em memória que é uma implementação do SQLite em Go. Ele é rápido, leve e fácil de usar. 

Trata-se de um banco de dados autônomo que armazena todo o banco de dados em um único arquivo e pode ser utilizado em memória, o que significa que não é necessário armazená-lo em disco, além de que, podermos utilizar este armazenamento de dados temporariamente durante a execução do programa, semelhantemente a um banco H2 utilizado em Java.


## Instalando e utilizando o `sqlite3`

Para utilizar o `sqlite3` em Go, é necessário importar o pacote `github.com/mattn/go-sqlite3` e, em seguida, criar e interagir com o banco de dados em memória conforme os passos abaixo:

1. Instalar o pacote `github.com/mattn/go-sqlite3`:

```bash
go get -u github.com/mattn/go-sqlite3
```

2. Certifique-se de estar em um diretório que já é um módulo Go, ou inicialize um novo módulo Go em seu diretório. Se está iniciando um novo projeto, vá para o diretório do seu projeto e execute:

```bash
go mod init nome_do_seu_modulo
```

Isso criará um arquivo `go.mod` no diretório raiz do seu projeto.

3. Em um arquivo por exemplo, `main.go` importar o pacote `github.com/mattn/go-sqlite3`:

```go
import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)
```

4. Criar um banco de dados em memória, ainda no arquivo `main.go` e no escopo de uma função principal por exemplo `main()`:

```go
db, err := sql.Open("sqlite3", ":memory:")
if err != nil {
    log.Fatal(err)
}
defer db.Close()
```

A partir deste ponto, é possível criar tabelas, inserir dados, executar consultas e tudo mais que faria com um banco de dados SQLite normal. Por exemplo:

```go
_, err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, email TEXT)")
if err != nil {
    log.Fatal(err)
}

_, err = db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Alice", "alice@example.com")
if err != nil {
    log.Fatal(err)
}

rows, err := db.Query("SELECT * FROM users")
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

for rows.Next() {
    var id int
	var name string
	var email string
	err := rows.Scan(&id, &name, &email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ID:", id, "Name:", name, "Email:", email)
}
```

5. Execute o programa com o comando:

```bash
go run main.go
```

O programa retornará ID: `1 Name: Alice Email: alice@example.com`


Neste exemplo, criamos um banco de dados em memória, criamos uma tabela de usuários, inserimos um usuário chamado "Alice", e então recuperamos e imprimimos todos os usuários na tabela.

O banco de dados em memória do SQLite em Go é uma opção conveniente para testes, desenvolvimento rápido ou qualquer outra situação em que você precise de um banco de dados temporário que não persista em disco.


## Banco de Dados em Memória vs Disco


### Em Memória

Conforme o código abaixo, estamos usando o driver SQLite `github.com/mattn/go-sqlite3` para abrir um banco de dados em memória. O argumento `:memory:` é especial para o SQLite e indica que queremos um banco de dados que existe apenas na memória RAM, não em um arquivo no sistema de arquivos.

```go
// Abrir o banco de dados em memória
var err error
db, err = sql.Open("sqlite3", ":memory:")
if err != nil {
    log.Fatal(err)
}
defer db.Close()
```

Ao executar o programa Go, um banco de dados SQLite será criado em memória, ou seja, existirá apenas enquanto o programa estiver em execução. Cada nova execução do programa resultará em um novo banco de dados em memória.


### Em Disco

Se a necessidade for criar um banco de dados em disco ao invés de em memória, podemos fazer assim:

```go
// Abrir o banco de dados em disco
db, err = sql.Open("sqlite3", "caminho/para/seu/banco.db")
if err != nil {
    log.Fatal(err)
}
defer db.Close()
```

Nesse caso, o banco de dados SQLite será criado em um arquivo `caminho/para/seu/banco.db` no sistema de arquivos. Se o arquivo não existir, ele será criado. Se já existir, o banco de dados será aberto.


## Criando uma REST API básica com acesso ao banco de dados

Todo o código abaixo pode ser escrito em um único arquivo `main.go.` 

Neste arquivo main.go, temos:

1. A definição da estrutura (struct) User que representa um usuário.
2. A abertura e fechamento do banco de dados SQLite em memória.
3. A criação da tabela users caso ela ainda não existir.
4. Três manipuladores de rotas para a API:
- `/users`: Busca todos os usuários.
- `/user`: Insere um novo usuário.
- `/user/{id}`: Busca um usuário por ID.
5. O servidor HTTP é iniciado na porta 8080 para servir essas rotas.

Certifique-se de que o projeto Go tenha o módulo habilitado (criando um arquivo go.mod na raiz do seu projeto), ou que você está trabalhando em um diretório dentro do GOPATH, para que o Go possa baixar e usar o pacote do SQLite corretamente.

Para gerar um módulo, execute o comando `go mod init nome_do_seu_modulo`.


```go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// Struct para representar um usuário
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *sql.DB

func main() {
	// Abrir o banco de dados em memória
	var err error
	db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Criar a tabela de usuários
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// Rotas da API
	http.HandleFunc("/users", handleUsers)   // Rota para buscar todos os usuários
	http.HandleFunc("/user", handleInsert)   // Rota para inserir um usuário
	http.HandleFunc("/user/", handleGetUser) // Rota para buscar um usuário por ID

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Manipulador para buscar todos os usuários
func handleUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

// Manipulador para inserir um usuário
func handleInsert(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.ID = int(lastID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Manipulador para buscar um usuário por ID
func handleGetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/user/"):]
	var user User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id=?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
```

### Testando a API

#### Inserir um Usuário (POST)

Para inserir um novo usuário na API, faremos uma requisição `POST` para o endpoint /user com os dados do usuário no corpo da requisição.

```bash
curl -X POST http://localhost:8080/user -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "john@example.com"}'
```

Resposta do POST (Inserir Usuário):

```json
{"ID":1,"Name":"John Doe","Email":"john@example.com"}
```

#### Buscar Todos os Usuários (GET)

Para buscar todos os usuários cadastrados na API, faremos uma requisição `GET` para o endpoint `/users`.

```bash
curl http://localhost:8080/users
```

Resposta do GET (Buscar Todos os Usuários):

```json
[{"ID":1,"Name":"John Doe","Email":"john@example.com"}]
```

#### Buscar um Usuário pelo ID (GET)

Para buscar um usuário específico pelo ID, faremos uma requisição `GET` para o endpoint `/user/{id}`. Substitua `{id}` pelo ID do usuário que deseja buscar.

```bash
curl http://localhost:8080/user/1
```

Neste exemplo, estamos buscando o usuário com ID igual a 1. Se o usuário com esse ID existir na base de dados, a resposta da API será semelhante a isso:

Exemplo de Resposta da Busca de um Usuário

```json
{"ID":1,"Name":"John Doe","Email":"john@example.com"}
```

Isso significa que o usuário com ID 1 foi encontrado na base de dados e as informações dele foram retornadas pela API.


Isso assume que a API está em execução localmente na porta 8080. Certifique-se de alterar os dados do corpo da requisição conforme necessário para os dados do usuário que você deseja inserir.

Se estiver usando uma ferramenta como o Postman ou Insomnia, os passos seriam semelhantes, mas você usaria a interface gráfica dessas ferramentas para configurar e enviar as requisições.


#### Liberando a Porta 8080

Caso a porta utilizada fique presa no processo, utilize o comando abaixo para liberar:

```bash
kill -9 $(lsof -t -i:8080)
```

## Conclusão

Nestr estudo, exploramos oo uso do SQLite3 em Go, aprendemos a criar um banco de dados em memória, criar tabelas, inserir e buscar registros. Através do pacote `database/sql` e `github.com/mattn/go-sqlite3`, construímos uma API simples para manipular dados em um banco SQLite. 

Essa abordagem oferece uma solução leve e eficiente para aplicações que precisam de um armazenamento local simples, sendo uma opção viável para muitos cenários de desenvolvimento. Com o Go, podemos facilmente integrar e interagir com bancos de dados, mantendo a simplicidade e a eficácia que a linguagem oferece.

