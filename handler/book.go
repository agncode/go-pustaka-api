package handler

import (
	"fmt"
	"net/http"

	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Agung G",
		"bio":  "Software Engineer",
		"age":  23,
	})
}

func HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Judul":   "Hello World",
		"isi":     "lorem ipsum dolor sit amet",
		"tanggal": "2023-10-01",
	})
}

func BooksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")

	ctx.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
	// ctx.JSON(http.StatusOK, gin.H{
	// 		"judul":"Belajar Golang",
	// 		"pengarang":"Agung G",
	// 		"penerbit":"Golang Books",
	// 	})
}

func QueryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	price := ctx.Query("price")

	ctx.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

func PostBooksHandler(ctx *gin.Context)  {
	var bookInput book.BookInput

	err := ctx.ShouldBindJSON(&bookInput)
	if  err != nil {

		errorMessages:=[]string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Field %s is not valid, condition: %s", e.Field(),e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
	
}