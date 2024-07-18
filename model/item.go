package model

type Item struct {
	ID          int64    `json:"id" excel:"id"`
	Name        string   `json:"name" excel:"name"`
	Price       float64  `json:"price" excel:"price"`
	Tags        []string `json:"tags" excel:"tags"`
	Description string   `json:"description" excel:"description"`
}

const ItemSheet = "item"
