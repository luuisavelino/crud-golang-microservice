package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/crud-golang-microservice-gestao-produtos/pkg/models"
	"net/http"
	"encoding/json"
)

func Products(c *gin.Context) {
	p := []models.Product{{
		Name:        "Camisa",
		Description: "Azul",
		Price:       10,
		Amount:      5,
	},
		{
			Name:        "Calça",
			Description: "Verde",
			Price:       120,
			Amount:      15,
		}}

	c.JSON(http.StatusOK, p)
}

func Product(c *gin.Context) {

	p := models.Product{
		Name:        "Camisa",
		Description: "Azul",
		Price:       10,
		Amount:      5,
	}

	c.JSON(http.StatusOK, p)
}

func DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("seq")
	fmt.Println("Id", id, "deletado com sucesso!")

	c.JSON(http.StatusOK, "Id "+id+" deletado com sucesso!")
}

func CreateProduct(c *gin.Context) {
	var product models.Product

	decoder := json.NewDecoder(c.Request.Body)
	decoder.Decode(&product)

	fmt.Println(product)

	c.JSON(http.StatusOK, "Item criado com sucesso!")
}

func UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("seq")
	fmt.Println("Id", id, "deletado com sucesso!")
	c.JSON(http.StatusOK, "Id "+id+" editado com sucesso!")
}
