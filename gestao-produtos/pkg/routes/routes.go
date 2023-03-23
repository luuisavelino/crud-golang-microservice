package routes

import (
	"github.com/luuisavelino/crud-golang-microservice-gestao-produtos/pkg/controllers"
	"net/http"
)

const apiVersion = "/api/v1"

func HandlerRequest() {
	http.HandleFunc(apiVersion + "/products", controllers.Products)

	http.ListenAndServe(":8081", nil)
}
