package main

import (
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
		
		bookRepository := book.NewRepository(db)//untuk membuat repository baru dengan nama bookRepository
		bookService := book.NewService(bookRepository)//untuk membuat service baru dengan nama bookService

		//Anggap ini adalah data yang ingin kita masukkan ke dalam database atau handler
		bookRequest := book.BookRequest{
			Title: "Belajar Jafa ke DUA",
			Description: "Buku ini adalah buku tentang belajar jafa ke DUA",
			Price: "48000",
		}
		bookService.Create(bookRequest)//untuk membuat buku baru dengan nama bookService



		books,err := bookRepository.FindAll()//untuk mengambil semua data buku dari database pustaka-api
		if err != nil {fmt.Print(err)}


		for _, book := range books {
			fmt.Println("==============================================")
			fmt.Println("Title:", book.Title)
			fmt.Println("Description:", book.Description)
			fmt.Println("Price:", book.Price)
			fmt.Println("==============================================")
		}

	router := gin.Default()//untuk membuat router baru

	v1 := router.Group("/v1")//untuk membuat grup router baru dengan prefix /v1

	v1.GET("/",handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query",handler.QueryHandler)
	v1.POST("/books",handler.PostBooksHandler)

	router.Run(":8888")

	//main
	//handler 
	//service
	//repository
	//ddb
	//mysql
}