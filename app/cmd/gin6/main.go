package main
// https://xminatolog.com/post/2518
// curl http://localhost:3000/2

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "fmt"
)

type snacks struct{
	ID		string `json:"id"`
	Title	string `json:"title"`
}

var snack1 = []snacks{
	{ID: "1", Title: "chocolate"},
	{ID: "2", Title: "pancake"},
}

func getJsonById(ctx *gin.Context) {
	// id, err := strconv.Atoi(ctx.Param("id"))
	id := ctx.Param("id")
	
	for _, k := range snack1 {
		if k.ID == id {
			// ctx.IndentedJSON(http.StatusOK, k)
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "ok", "ID": k.ID, "Title": k.Title})
			// ctx.IndentedJSON(http.StatusOK, gin.H{"Hello! ID: %s\n": "This ID's title: %s", k.ID, k.Title})
			// ctx.IndentedJSON(http.StatusOK, gin.H{"Hello! ID: ", k.ID, k.Title})
			// ctx.String(http.StatusOK, "Hello! ID: %s\nThis ID's title: %s", k.ID, k.Title)
			// ctx.String(http.StatusOK, "Hello! title: %s",k.Title)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
}


func openMainPage(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)

	// ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found index.html"})
}

func main(){
	// fmt.Println("Hello!")
	router := gin.Default()
	// router.Get("/")
	router.LoadHTMLGlob("./htmls/*.html")
	router.GET("/", openMainPage)
	router.GET("/:id", getJsonById)
	router.Run()
}