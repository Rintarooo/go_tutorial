package main
// https://xminatolog.com/post/2518
// curl http://localhost:3000/2

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type Result struct {
	TrackName   string `json:"trackName"`
	ArtistName  string `json:"artistName"`
}

type Response struct {
	ResultCount int      `json:"resultCount"`
	Results     []Result `json:"results"`
}

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

func postSearchItunesApi(ctx *gin.Context){
	// http://127.0.0.1:3000/api/itunes/search/punpee
	artist := ctx.PostForm("artist")
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
	// ctx.Data(http.StatusOK, "application/json", body)
	// ctx.Data(http.StatusOK, "application/json", body.results)
	var response Response
  err = json.Unmarshal(body, &response)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, result := range response.Results {
		// fmt.Println("Track Name:", result.TrackName)
		// fmt.Println("Artist Name:", result.ArtistName)
		ctx.IndentedJSON(500, gin.H{"Artist Name:" : result.ArtistName, "Track Name:" : result.TrackName})
	}
}

func main(){
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*.html")
	router.GET("/", openMainPage)
	router.GET("/api/mock", getMockApi)
	// http://127.0.0.1:3000/api/itunes/search/punpee
	router.POST("/api/itunes/search/", postSearchItunesApi)
	router.Run()
}