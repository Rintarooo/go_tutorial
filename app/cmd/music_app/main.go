package main

import (
	"github.com/gin-gonic/gin"
)

func openMainPage(ctx *gin.Context){
	ctx.HTML(200, "index.html", nil)	
}

func main(){
	router := gin.Default()
	router.LoadHTMLGlob("./htmls/*.html")
	router.GET("/", openMainPage)
	router.Run()
}