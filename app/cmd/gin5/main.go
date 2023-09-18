// https://zenn.dev/satumahayato010/articles/dd4a5c68ca6d67
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, books)
	// ctx.HTML(200, "index.html", nil)
}

func main() {
	router := gin.Default()

	

	router.LoadHTMLGlob("templates/*.html")

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.HTML(200, "index.html", nil)
	// })

	router.GET("/books", getBooks)


	router.Run()

}