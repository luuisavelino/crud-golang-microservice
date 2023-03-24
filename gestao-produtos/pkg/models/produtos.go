package models

import (
	"github.com/luuisavelino/crud-golang-microservice-gestao-produtos/pkg/database"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Amount      int     `json:"amount" binding:"required"`
}


func BuscaTodosOsProdutos() []Product {
	db := database.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	produtos := []Product{}

	for selectDeTodosOsProdutos.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = selectDeTodosOsProdutos.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func CriarNovoProduto(product Product) {
	db := database.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(name, description, price, amount) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(product.Name, product.Description, product.Price, product.Amount)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := database.ConectaComBancoDeDados()

	deletarOProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOProduto.Exec(id)
	defer db.Close()
}

func BuscaProduto(id string) Product {
	db := database.ConectaComBancoDeDados()

	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produto := Product{}

	for produtoDoBanco.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = produtoDoBanco.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Name = name
		produto.Description = description
		produto.Price = price
		produto.Amount = amount
	}

	defer db.Close()
	return produto
}

func AtualizaProduto(product Product) {
	db := database.ConectaComBancoDeDados()

	AtualizaProduto, err := db.Prepare("Update produtos set name=$1, description=$2, price=$3, amount=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	AtualizaProduto.Exec(product.Name, product.Description, product.Price, product.Amount)
	defer db.Close()
}
