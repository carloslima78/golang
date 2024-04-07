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
