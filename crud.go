package main

import (
	initializers "goCRUD/Initializers"
	"goCRUD/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	// LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
