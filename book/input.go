package book

import "encoding/json"

type BookInput struct {// BookInput represents the input for creating a new book
	Title string      `json:"title" binding:"required"` //untuk mengharuskan field ini diisi
	Price json.Number `json:"price" binding:"required,number"`//supaya field ini diisi dan harus berupa number
}