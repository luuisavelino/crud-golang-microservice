package models

type Product struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Amount      int     `json:"amount" binding:"required"`
}
