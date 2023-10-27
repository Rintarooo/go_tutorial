package main
// https://xminatolog.com/post/2518
// curl http://localhost:3000/2

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
)

func openMainPage(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}

func getItunesApi(ctx *gin.Context){
	res, err := http.Get("http://httpbin.org/get")
	if err != nil{
		ctx.IndentedJSON(500, gin.H{"error message: ": err.Error()})
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
			ctx.IndentedJSON(500, gin.H{"message": err.Error(),})
			return
	}

	ctx.String(200, string(body))

}

func main(){
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*.html")
	router.GET("/", openMainPage)
	router.GET("/api_songs", getItunesApi)
	router.Run()
}