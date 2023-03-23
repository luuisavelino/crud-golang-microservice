package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiVersion = "/api/v1"

type Product struct {
	Name        string
	Description string
	Price       string
	Amount      int
}

type Actions interface {
	SearchAll()
	Get()
	Delete()
	AddNew()
	Update()
}

func ApiConsume(method, apiKey, endpoint string, reqBody io.Reader) (io.Reader, error) {
	client := &http.Client{}
	req, err := http.NewRequest(
		method, fmt.Sprintf(apiVersion+BaseURL+endpoint), reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-key", apiKey)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return res.Body, nil
}

func (p *Products) SearchAll() ([]Product, error) {
	reqBody := bytes.NewBufferString(``)

	Body, err := ApiConsume("GET", p.apiKey, "/", reqBody)
	if err != nil {
		return nil, err
	}

	var produtos []Product
	if err := json.NewDecoder(Body).Decode(&produtos); err != nil {
		return nil, err
	}

	return produtos, nil
}

func (p *Products) Get(id string) (Product, error) {
	reqBody := bytes.NewBufferString(``)

	Body, err := ApiConsume("GET", p.apiKey, "/", reqBody)
	if err != nil {
		return Product{}, err
	}

	var produto Product
	if err := json.NewDecoder(Body).Decode(&produto); err != nil {
		return Product{}, err
	}

	return produto, nil
}

func (p *Products) Delete(id string) {
	fmt.Println("Produto deletado")
}

func (p *Products) AddNew() {
	fmt.Println("Novo produto criado")
}

func (p *Products) Update() {
	fmt.Println("Produto atualizado")
}
