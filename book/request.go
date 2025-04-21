package book

import "encoding/json"

type BookRequest struct {// BookInput represents the input for creating a new book
	Title string      `json:"title" binding:"required"` //untuk mengharuskan field ini diisi
	Price json.Number `json:"price" binding:"required,number"`//supaya field ini diisi dan harus berupa number
	Description string      `json:"description"` //untuk menampung deskripsi buku
	Rating	int `json:"rating"` //untuk menampung rating buku
	Discount int `json:"discount"` //untuk menampung diskon buku
}