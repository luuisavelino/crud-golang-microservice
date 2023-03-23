package routes

import (
	"github.com/luuisavelino/crud-golang-microservice-frontend/pkg/controllers"
	"net/http"
)

const apiVersion = "/api/v1"

func HandlerRequest() {
	http.HandleFunc(apiVersion + "/", controllers.Index)
	http.HandleFunc(apiVersion + "/new", controllers.New)
	http.HandleFunc(apiVersion + "/insert", controllers.Insert)
	http.HandleFunc(apiVersion + "/delete", controllers.Delete)
	http.HandleFunc(apiVersion + "/edit", controllers.Edit)
	http.HandleFunc(apiVersion + "/update", controllers.Update)

	http.ListenAndServe(":8080", nil)
}
