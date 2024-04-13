# Copier: Simplificando a Cópia de Structs em Golang

A manipulação de dados entre diferentes estruturas (`structs`) pode ser uma tarefa tediosa em linguagens como Go. Para simplificar esse processo, o pacote `Copier` oferece uma solução elegante e direta. 
Vamos explorar como isso pode ser feito.


## O que é o Copier?

O `Copier` é um pacote para Go que permite copiar dados de um tipo de struct para outro de forma simples e eficiente. Ele evita a necessidade de escrever código manualmente para atribuir campo por campo, facilitando o trabalho com estruturas de dados complexas.


### Documentação Oficial do Copier

A documentação oficial do `Copier` pode ser encontrada no [GitHub: Copier](https://github.com/jinzhu/copier)
. Lá podemos encontrar mais detalhes sobre como usar o pacote, exemplos adicionais e informações sobre as opções disponíveis para a cópia de dados entre structs em Go.


## Mão na massa

Vamos considerar um exemplo simples onde temos duas estruturas: `UserRequest` e `UserDomain`, cada uma representando diferentes estados de um usuário. A `UserRequest` é o objeto de entrada, enquanto `UserDomain` é o objeto de domínio com mais detalhes.


### Instalando o Copier

Antes de começar, certifique-se de ter o `Copier`instalado em seu ambiente Go. Você pode obtê-lo facilmente utilizando o comando:

```bash
go get github.com/jinzhu/copier
```

### Tag copier

É importante notar a presença da tag `copier` em cada campo das structs que compõem este exemplo. Esta tag é necessária para o funcionamento do `Copier` e é utilizada para indicar quais campos devem ser copiados e como eles devem ser mapeados. 

Vale observar que o `Copier` considera o nome de campo informado na tag para realizar o mapeamento dos campos, e desconsiera o nome do campo de origem na struct. Portanto, se o nome definido na tag for diferente do nome do campo na struct, o `Copier` consierará apenas o nome definido na tag.

Por exemplo, na struct UserRequest:

O campo ID possui a tag `copier:"ID"`, indicando que ele deve ser copiado para o campo ID da struct `UserDomain`.
O mesmo se aplica aos demais campos, como `Name`, `Email`, `Password` e `Address`.


### Estrutura de Dados: UserRequest

Aqui está a definição da struct UserRequest, que representa um objeto de requisição de um usuário:

```go
type UserRequest struct {
	ID       string         `copier:"ID"`
	Name     string         `copier:"Name"`
	Email    string         `copier:"Email"`
	Password string         `copier:"Password"`
	Address  AddressRequest `copier:"Address"`
}
```

A struct `UserRequest` possui campos como `ID`, `Name`, `Email`, `Password` e `Address`, que são os dados de um usuário que queremos copiar para outra struct.


### Estrutura de Dados: AddressRequest

Aqui está a definição da struct `AddressRequest`, que representa um objeto de requisição de endereço de um usuário:

```go
type AddressRequest struct {
	Street string `copier:"Street"`
	City   string `copier:"City"`
	State  string `copier:"State"`
	Zip    string `copier:"Zip"`
}
```

A struct `AddressRequest` possui campos como `Street`, `City`, `State` e `Zip`, que são os dados de endereço de um usuário.


### Estrutura de Dados: UserDomain

Aqui está a definição da struct `UserDomain`, que representa um objeto de domínio de um usuário:

```go
type UserDomain struct {
	ID       string        `copier:"ID"`
	Name     string        `copier:"Name"`
	Email    string        `copier:"Email"`
	Password string        `copier:"Password"`
	Address  AddressDomain `copier:"Address"`
}           `
```

A struct `UserDomain` possui campos como `ID`, `Name`, `Email`, `Password` e `Address`, que são os dados completos de um usuário, incluindo seu endereço.


### Estrutura de Dados: AddressDomain

Aqui está a definição da struct `AddressDomain`, que representa um objeto de domínio de um endereço de usuário:

```go
type AddressDomain struct {
	Street string `copier:"Street"`
	City   string `copier:"City"`
	State  string `copier:"State"`
	Zip    string `copier:"Zip"`
}
```

A struct `AddressDomain` possui campos como `Street`, `City`, `State` e `Zip`, que são os dados completos de endereço de um usuário.


### Função Main: Testando a conversão das structs

Aqui está a função `main()` onde realizamos a cópia dos dados de UserRequest para UserDomain usando o `Copier`:

```go
func main() {
	// Criando um objeto de requisição de usuário
	userRequest := UserRequest{
		ID:       "1",
		Name:     "Carlos Soares",
		Email:    "carlos.soares@example.com",
		Password: "123456",
		Address: AddressRequest{
			Street: "Rua A",
			City:   "São Paulo",
			State:  "SP",
			Zip:    "01234-567",
		},
	}

	// Criando um objeto de domínio vazio para receber a cópia do objeto de requisição
	userDomainCopied := UserDomain{}

	// Copiando o objeto de requisição para o objeto de domínio
	err := copier.Copy(&userDomainCopied, &userRequest)

	// Verificando se ocorreu algum erro durante a cópia
	if err != nil {
		fmt.Println(err)
		return
	}

	// Imprimindo o objeto de domínio copiado
	fmt.Printf("Converted UserRequest to UserDomain: %#v \n", userDomainCopied)
}
```

Neste exemplo, a linha `err := copier.Copy(&userDomainCopied, &userRequest)` é crucial. Aqui, utilizamos o método `copier.Copy` passando como argumentos os objetos `userDomainCopied` que receberá a cópia, e `userRequest` que contém os dados de entrada e será a origem dos dados. 

O `Copier` então mapeia automaticamente os campos das structs `UserRequest` e `UserDomain` com base nas tags `copier` e realiza a cópia dos dados. Se ocorrer qualquer erro durante este processo, ele será capturado pela variável `err`.


#### Executando o código

Para executar o código, basta abrir um terminal na pasta onde está o arquivo `main.go` e executar o comando abaixo:

```bash
go run main.go
```

Isso compilará e executará o programa. A saída no terminal deve apresentar o objeto `userDomainCopied` com os dados copiados do `userRequest` conforme abaixo:


```bash
Converted UserRequest to UserDomain: main.UserDomain{ID:"1", Name:"Carlos Soares", Email:"carlos.soares@example.com", Password:"123456", Address:main.AddressDomain{Street:"Rua A", City:"São Paulo", State:"SP", Zip:"01234-567"}} 
```


## Conclusão

O pacote `Copier` para Go simplifica o processo de cópia entre structs, economizando tempo e reduzindo a chance de erros, além de oferecer uma sintaxe limpa e eficiente. Smplifique sua vida de desenvolvedor!