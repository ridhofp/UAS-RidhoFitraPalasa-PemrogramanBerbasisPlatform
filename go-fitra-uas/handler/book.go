package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go-fitra-uas/book"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":  "ridho",
		"title": "hello ido",
	})

}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":  "cek",
		"title": "cekk",
	})

}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"id": id})

}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}
func PostBooksHandler(c *gin.Context) {
	//title, price

	var bookInput book.BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}
