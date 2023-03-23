package controllers

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Amount      int
}

func Products(w http.ResponseWriter, r *http.Request) {
	p := Product{
		Name: "Camisa",
		Description: "Azul",
		Price: 10,
		Amount: 5,
	}
	
	json.NewEncoder(w).Encode(p)
}

