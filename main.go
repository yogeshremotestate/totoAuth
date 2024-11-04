package main

import (
	"LearnGo-todoAuth/controllers"
	"LearnGo-todoAuth/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.POST("/notes/", controllers.NoteCreate)
	r.GET("/notes/", controllers.GetAllNote)
	r.GET("/notes/:id", controllers.GetNote)
	r.PUT("/notes/:id", controllers.UpdateNote)
	r.DELETE("/notes/:id", controllers.DeleteNote)
	r.Run()
}
