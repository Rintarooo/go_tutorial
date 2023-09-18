// https://zenn.dev/ajapa/articles/6471ac0c612fda
package main

import (
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
)

func nl2br(text string) template.HTML {
	return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br />", -1))
}

func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"nl2br": nl2br,
	})
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})
	router.POST("/result", func(ctx *gin.Context) {
		result, _ := ctx.GetPostForm("result")
		ctx.HTML(200, "result.html", result)
	})

	router.Run()
}
