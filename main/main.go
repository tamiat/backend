package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type content struct {
	/*
		Struct tags such as json:"id" specify what a field's name should be when the struct's contents are serialized into JSON.
		without them, the JSON would use the struct's capitalized field names - a style not as common in JSON
	*/
	ID    string `json:"id"`
	Title string `json:"title"`
	Details string `json:"details"`
}

// getContents responds with the list of all contents as JSON.
func getContents(c *gin.Context) {
	//serialize the struct into JSON and add it to the response.
    c.IndentedJSON(http.StatusOK, contents)
}

//Some arbitrary data
var contents = []content{
	{ID:"0",Title:"Hellooooo World!",Details:"Welcome to Tamiat CMS!"},
	{ID:"1",Title:"Hellooooo monde!",Details:"Bienvenue à Tamiat CMS!"},
	{ID:"2",Title:"Hola Mundo",Details:"Bienvenida a Tamiat CMS!"},
}

func main() {
	//Initialzing a Gin router using Default
    router := gin.Default()
    router.GET("/contents", getContents)
    router.Run("localhost:8080")
}

