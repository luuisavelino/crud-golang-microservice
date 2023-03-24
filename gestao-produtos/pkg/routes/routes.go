package routes

import (
	"github.com/luuisavelino/crud-golang-microservice-gestao-produtos/pkg/controllers"
	"github.com/gin-gonic/gin"
)

const apiVersion = "/api/v1"

func HandlerRequest() {
	router := gin.Default()

	router.GET(apiVersion + "/products", controllers.Products)
	router.GET(apiVersion + "/products/:id", controllers.Product)
	router.POST(apiVersion + "/products", controllers.CreateProduct)
	router.PATCH(apiVersion + "/products/:id", controllers.UpdateProduct)
	router.DELETE(apiVersion + "/products/:id", controllers.DeleteProduct)

	router.Run(":8080")
}
