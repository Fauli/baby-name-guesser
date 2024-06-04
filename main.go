package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "sbebe.ch/baby-name-guesser/docs"
	"sbebe.ch/baby-name-guesser/pkg/controller"
)

const basePath = "/api"

func main() {
	docs.SwaggerInfo.BasePath = basePath

	// Create a new Gin router
	router := gin.Default()

	apiGroup := router.Group(basePath)
	nameGroup := apiGroup.Group("/names")
	votersGroup := apiGroup.Group("/voters")
	votesGroup := apiGroup.Group("/votes")

	c := controller.NewController()

	// Define your API routes
	nameGroup.GET("", c.GetAllBabyNames)
	nameGroup.POST("", c.AddBabyNames)
	nameGroup.DELETE("/:name", c.DeleteBabyName)

	votersGroup.POST("", c.AddNewVoter)
	votersGroup.POST("/login", c.LoginVoter)
	votersGroup.GET("/:email", c.GetVoter)
	votersGroup.DELETE("/:email", c.DeleteVoter)

	votesGroup.POST("", c.AddVotes)
	votesGroup.GET("", c.GetAllVotes)
	votesGroup.GET("/voters", c.GetVotesPerVoters)

	// Serve the Swagger documentation
	apiGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Start the server
	router.Run(":8080")
}
