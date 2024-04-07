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
