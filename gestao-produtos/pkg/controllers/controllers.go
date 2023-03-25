package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	var product models.Product

	c.Request.ParseMultipartForm(1000)
	for key, value := range c.Request.PostForm {
		switch key {
		case "name":
			product.Name = value[0]	
		case "description":
			product.Description = value[0]
		case "price":
			price, err := strconv.ParseFloat(value[0], 64)
			if err != nil {
				log.Println(err)
			}
			product.Price = price
		case "amount":
			amount, err := strconv.Atoi(value[0])
			if err != nil {
				log.Println(err)
			}
			product.Amount = amount
		}
	}

	models.CriarNovoProduto(product)

	c.JSON(http.StatusOK, "Item criado com sucesso!")
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	id, err := strconv.Atoi(c.Params.ByName("id"))	
	if err != nil {
		c.JSON(http.StatusBadRequest, "Id do produto inválido!")
	}

	c.Request.ParseMultipartForm(1000)
	for key, value := range c.Request.PostForm {
		switch key {
		case "name":
			product.Name = value[0]	
		case "description":
			product.Description = value[0]
		case "price":
			price, err := strconv.ParseFloat(value[0], 64)
			if err != nil {
				log.Println(err)
			}
			product.Price = price
		case "amount":
			amount, err := strconv.Atoi(value[0])
			if err != nil {
				log.Println(err)
			}
			product.Amount = amount
		}
	}

	product.Id = id

	models.AtualizaProduto(product)

	c.JSON(http.StatusOK, "Id "+string(id)+" editado com sucesso!")
}
