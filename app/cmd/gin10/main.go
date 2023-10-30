package main
// https://xminatolog.com/post/2518
// curl http://localhost:3000/2

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"fmt"
)

func openMainPage(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}

func getMockApi(ctx *gin.Context){
	res, err := http.Get("http://httpbin.org/get")
	// res, err := http.Get("https://itunes.apple.com/search")
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

func getSearchItunesApi(ctx *gin.Context){
	// http://127.0.0.1:3000/api/itunes/search/punpee
	artist := ctx.Param("artist")
	res, err := http.Get(fmt.Sprintf("https://itunes.apple.com/search?term=%s&entity=song", artist))
	if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}
	ctx.Data(http.StatusOK, "application/json", body)
}

func main(){
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*.html")
	router.GET("/", openMainPage)
	router.GET("/api/mock", getMockApi)
	// http://127.0.0.1:3000/api/itunes/search/punpee
	router.GET("/api/itunes/search/:artist", getSearchItunesApi)
	router.Run()
}