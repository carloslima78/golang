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
