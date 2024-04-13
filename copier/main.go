package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

// UserDomain representa um objeto de domínio de um usuário
type UserDomain struct {
	ID       string        `copier:"ID"`
	Name     string        `copier:"Name"`
	Email    string        `copier:"Email"`
	Password string        `copier:"Password"`
	Address  AddressDomain `copier:"Address"`
}

// AddressDomain representa um objeto de domínio de um endereço de um usuário
type AddressDomain struct {
	Street string `copier:"Street"`
	City   string `copier:"City"`
	State  string `copier:"State"`
	Zip    string `copier:"Zip"`
}

// UserRequest representa um objeto de requisição de um usuário
type UserRequest struct {
	ID       string         `copier:"ID"`
	Name     string         `copier:"Name"`
	Email    string         `copier:"Email"`
	Password string         `copier:"Password"`
	Address  AddressRequest `copier:"Address"`
}

// AddressRequest representa um objeto de requisição de um endereço de usuário
type AddressRequest struct {
	Street string `copier:"Street"`
	City   string `copier:"City"`
	State  string `copier:"State"`
	Zip    string `copier:"Zip"`
}

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
