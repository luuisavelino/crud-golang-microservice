package models

import "os"

var BaseURL = os.Getenv("BASE_URL")

type Products struct {
	apiKey string
}

func New(apiKey string) *Products {
	return &Products{apiKey: apiKey}
}
