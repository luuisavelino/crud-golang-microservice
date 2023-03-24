package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/crud-golang-microservice-gestao-produtos/pkg/models"
	"net/http"
	"encoding/json"
)

func Products(c *gin.Context) {
	products := models.BuscaTodosOsProdutos()
	c.JSON(http.StatusOK, products)
}

func Product(c *gin.Context) {
	id := c.Params.ByName("id")
	product := models.BuscaProduto(id)
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	models.DeletaProduto(id)
	c.JSON(http.StatusOK, "Id "+id+" deletado com sucesso!")
}

func CreateProduct(c *gin.Context) {
	var product models.Product

	decoder := json.NewDecoder(c.Request.Body)
	fmt.Println("decoder")
	decoder.Decode(&product)

	fmt.Println(product)

	models.CriarNovoProduto(product)

	c.JSON(http.StatusOK, "Item criado com sucesso!")
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Params.ByName("id")

	decoder := json.NewDecoder(c.Request.Body)
	decoder.Decode(&product)

	fmt.Println(product)

	models.AtualizaProduto(product)
	c.JSON(http.StatusOK, "Id "+id+" editado com sucesso!")
}
