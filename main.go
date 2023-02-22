package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
	Welcome string `json:"welcome"`
}

type Recipe struct {
	Name string `json:"name"`
	Description string `json:"description"`
	IsDifficult bool `json:"is_difficult"`
}

func main() {
	fmt.Println("Hello World")

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		fmt.Println("Calling ping function")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"welcome": "you are welcome",
		})
	})

	router.POST("/recipe", func(ctx *gin.Context) {
		var requestRecipe Recipe

		err := ctx.ShouldBindJSON(&requestRecipe)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message":"error occured",
			})

			return
		}

		fmt.Println("Recipe recieved: ", requestRecipe)

		ctx.JSON(200, gin.H{
			"message": "request was successful",
		})


	})

	router.POST("/ping", func(ctx *gin.Context) {
		var requestBody Message

		err := ctx.ShouldBindJSON(&requestBody)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "invalid data type",
			})

			return
		}

		fmt.Println("Message received: ", requestBody)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "successfully received message",
		})
	})

	router.Run()
}

// add post request
// add json types
// use functions

// store values in a slice