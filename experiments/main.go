package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Model of book entity
type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// Mock database for learning purpose
var books = []book{
	{ID: "1", Title: "Potop", Author: "Henryk Sienkiewicz", Quantity: 2},
	{ID: "2", Title: "Wesele", Author: "Stanisław Wyspiański", Quantity: 5},
	{ID: "3", Title: "Krzyżacy", Author: "Henryk Sienkiewicz", Quantity: 2},
	{ID: "4", Title: "Ferdydurke", Author: "Witold Gombrowicz", Quantity: 4},
}

// Implementation of API
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)

}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("Book not found")
}

func getBooksByAuthor(author string) ([]*book, error) {
	var booksByAuthor []*book
	for i, b := range books {
		if b.Author == author {
			booksByAuthor = append(booksByAuthor, &books[i])
		}
	}

	if len(booksByAuthor) == 0 {
		return nil, errors.New("Books not found.")
	}

	return booksByAuthor, nil
}

func booksByAuthor(c *gin.Context) {
	author, ok := c.GetQuery("author")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing author query parameter."})
		return
	}

	authors, err := getBooksByAuthor(author)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Books not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, authors)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book with this id doesn't exist."})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)

}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)

}

func createBook(c *gin.Context) {
	var newBook book

	err := c.BindJSON(&newBook)
	if err != nil {
		fmt.Println("Error binding newBook values!")
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.GET("/books/by_author", booksByAuthor)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")

}
