package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "sbebe.ch/baby-name-guesser/docs"
	"sbebe.ch/baby-name-guesser/pkg/controller"
	middleware "sbebe.ch/baby-name-guesser/pkg/middlware"

	"github.com/gin-contrib/cors"
)

const basePath = "/api"

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	sessionKey = []byte("tandaradeiIstDasGoldeneVomEiGut!")
	// https://gowebexamples.com/sessions/
	store = sessions.NewCookieStore(sessionKey)
)

func main() {
	docs.SwaggerInfo.BasePath = basePath

	// Create a new Gin router
	router := gin.Default()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// validate if users are logged in
	router.Use(middleware.ValidateSession(store))

	// allow all cores traffic for now
	// TODO: [franz] fix this to only allow the frontend
	// https://github.com/gin-contrib/cors
	router.Use(cors.Default())

	apiGroup := router.Group(basePath)
	nameGroup := apiGroup.Group("/names")
	votersGroup := apiGroup.Group("/voters")
	votesGroup := apiGroup.Group("/votes")

	c := controller.NewController(store)

	// Define your API routes
	nameGroup.GET("", c.GetAllBabyNames)
	nameGroup.POST("", c.AddBabyNames)
	nameGroup.DELETE("/:name", c.DeleteBabyName)

	votersGroup.POST("", c.AddNewVoter)
	votersGroup.POST("/login", c.LoginVoter)
	votersGroup.GET("/:email", c.GetVoter)
	votersGroup.DELETE("/:email", c.DeleteVoter)
	votersGroup.POST("/logout", c.LogoutVoter)

	votesGroup.GET("", c.GetAllVotes)
	votesGroup.POST("", c.AddVotes)
	votesGroup.GET("/voters", c.GetVotesPerVoters)
	votesGroup.GET("/top", c.GetTopVotes)

	// Serve the Swagger documentation
	apiGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Start the server
	router.Run(":8080")
}
