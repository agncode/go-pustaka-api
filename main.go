package main

import (
	"encoding/json"
	"fmt"
	"log"

	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	  if err != nil {
		  log.Println("Failed to connect database")
		}
		
		db.AutoMigrate(&book.Book{})//untuk migrasi tabel book ke database pustaka-api

		//CRUD
		//==========================
		// Create Data
		//==========================
		// book := book.Book{}
		// book.Title=       "Atomic Habits"
		// book.Description= "Belajar habit atau kebiasaan menggunakan buku ini"
		// book.Price=       120000
		// book.Discount= 	25
		// book.Rating =   4

		// err = db.Create(&book).Error
		// if err != nil {
		// 	fmt.Println("=======================================")
		// 	fmt.Println("Failed to create book record")
		// 	fmt.Println("=======================================")
		// } 

		var books []book.Book //untuk menampung data dari tabel book

		err = db.Debug().Where("title = ?","Belajar Golang").Find(&books).Error //untuk mengisi variabel books dengan data dari tabel book
		if err != nil {
			fmt.Println("=======================================")
			fmt.Println("Failed to find book record")
			fmt.Println("=======================================")	
		}

		for _, book := range books {//untuk menampilkan data book
			fmt.Println("=======================================")
			fmt.Println("Title: ", book.Title)
			fmt.Println("Description: ", book.Description)
			fmt.Printf("Book Object %v\n", book) //untuk menampilkan data book
		}

		//var book book.Book //untuk menampung data dari tabel book

		// err= db.Debug().First(&book,4).Error //untuk mengisi variabel book dengan data pertama dari tabel book
		// if err != nil {
		// 	fmt.Println("=======================================")
		// 	fmt.Println("Failed to find book record")
		// 	fmt.Println("=======================================")
		// }

		// fmt.Println("Title: ", book.Title)
		// fmt.Println("Description: ", book.Description)
		// fmt.Printf("Book Object %v ", book.Price) //untuk menampilkan data book

		
		
	router := gin.Default()//untuk membuat router baru

	v1 := router.Group("/v1")//untuk membuat grup router baru dengan prefix /v1

	v1.GET("/",handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query",handler.QueryHandler)
	v1.POST("/books",handler.PostBooksHandler)

	router.Run(":8888")
}
type BookInput struct{
	Title string 		`json:"title" binding:"required"`
	Price json.Number	`json:"price" binding:"required,number"`
}
