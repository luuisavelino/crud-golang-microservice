package models

import "fmt"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	var produtos []Produto
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	fmt.Println("Novo produto criado")
}

func DeletaProduto(id string) {
	fmt.Println("Produto deletado")
}

func EditaProduto(id string) Produto {
	var produtoParaAtualizar Produto
	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	fmt.Println("Produto atualizado")
}
