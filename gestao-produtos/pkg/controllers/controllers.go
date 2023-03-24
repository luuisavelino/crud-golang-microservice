package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/crud-golang-microservice-gestao-produtos/pkg/models"
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

	c.Request.ParseMultipartForm(1000)
	for key, value := range c.Request.PostForm {
		fmt.Println(key, value)
	}

	c.JSON(http.StatusOK, "Item criado com sucesso!")
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Params.ByName("id")

	if err := c.BindJSON(&product); err != nil {
		panic(err)
	}

	models.AtualizaProduto(product)
	c.JSON(http.StatusOK, "Id "+id+" editado com sucesso!")
}
