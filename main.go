package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books []Book

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", postBooks)
	router.GET("/books/:id", getBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)
	router.Run("localhost:8080")
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBooks(c *gin.Context) {
	var book Book
	err := c.BindJSON(&book)
	if err != nil {
		return
	}
	for _, b := range books {
		if book.ID == b.ID {
			c.Status(http.StatusBadRequest)
			return
		}
	}
	books = append(books, book)
	c.Status(http.StatusCreated)
}

func getBook(c *gin.Context) {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	for _, b := range books {
		if id == b.ID {
			c.IndentedJSON(http.StatusOK, b)
		}
	}
	c.Status(http.StatusNotFound)
}

func deleteBook(c *gin.Context) {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	for i, b := range books {
		if id == b.ID {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	c.Status(http.StatusNoContent)
}

func updateBook(c *gin.Context) {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	var updatedBook Book
	if err := c.BindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	for idx, b := range books {
		if b.ID == id {
			books[idx] = updatedBook
			c.Status(http.StatusAccepted)
			return
		}
	}
	c.Status(http.StatusNotFound)
	// nb := reflect.ValueOf(updatedBook)
	// typeOfnb := nb.Type()
	// for _, b := range books {
	// 	if b.ID == id {
	// 		for i := 0; i < nb.NumField(); i++ {
	// 			fmt.Printf("Field: %v\tValue: %v\n", typeOfnb.Field(i), nb.Field(i).Interface())
	// 		}
	// 	}
	// }
}
