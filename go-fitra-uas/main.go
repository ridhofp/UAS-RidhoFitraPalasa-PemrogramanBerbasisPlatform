package main

import (
	"fmt"
	"go-fitra-uas/book"
	"go-fitra-uas/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/go-fitra-uas?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	db.AutoMigrate(&book.Book{})
	//crud

	//--------create data-----------
	book := book.Book{}
	book.Title = "Fiersa Besari"
	book.Price = 70000
	book.Rating = 5
	book.Description = "mari beli produk buku ini"

	err = db.Create(&book).Error
	if err != nil {
		fmt.Println("=======================")
		fmt.Println("============error=========")
		fmt.Println("=======================")

	}

	//-------------------read data---------------------
	/*var book book.Book
	err = db.First(&book).Error
	if err != nil {
		fmt.Println("=======================")
		fmt.Println("===error finding book==")
		fmt.Println("=======================")
	}
	fmt.Println("Title :", book.Title)
	fmt.Println("book object : %v", book)*/

	//----------------------update data---------------

	/*var book book.Book
	err = db.Debug().Where("id = ?", 1).First(&book).Error
	if err != nil {
		fmt.Println("=======================")
		fmt.Println("===error update book==")
		fmt.Println("=======================")
	}
	book.Title = "tereliye (revisi)"
	err = db.Save(&book).Error*/

	//-----------------------delete data-------------
	/*var book book.Book
	err = db.Debug().Where("id = ?", 1).First(&book).Error
	if err != nil {
		fmt.Println("=======================")
		fmt.Println("===error delete book==")
		fmt.Println("=======================")
	}
	err = db.Delete(&book).Error*/

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8888")
}
