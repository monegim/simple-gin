package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books []Book

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", postBooks)
	router.Run("localhost:8080")
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBooks(c *gin.Context) {
	var newAlbum Book
	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}
	books = append(books, newAlbum)
	c.Status(http.StatusCreated)
}
