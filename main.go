package main

import (
	"fmt"

	"example.com/go-psql-es/controllers"
	"example.com/go-psql-es/database"

	//"example.com/go-psql-es/scripts"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting application ...")
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	database.DatabaseConnection()
	database.ElasticSearchConnection()
	//scripts.IngestBulkData()
	r := gin.Default()
	r.GET("/users/:id", controllers.ReadUser)
	r.GET("/users", controllers.ReadUsers)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	//r.DELETE("/users/:id", controllers.DeleteBook)
	r.GET("/projects/:id", controllers.ReadProject)
	r.GET("/projects", controllers.ReadProjects)
	r.POST("/projects", controllers.CreateProject)
	r.PUT("/projects/:id", controllers.UpdateProject)
	r.PATCH("/projects/:id", controllers.PatchProject)
	//user projects
	r.GET("/userprojects/:id", controllers.ReadUserProject)
	r.GET("/userprojects", controllers.ReadUserProjects)
	r.POST("/userprojects", controllers.CreateUserProject)
	r.PUT("/userprojects/:id", controllers.UpdateUserProject)
	r.PATCH("/userprojects/:id", controllers.PatchUserProject)
	r.GET("/userprojects/ingest/:id", controllers.Ingest)
	//hastags
	r.GET("/hashtags/:id", controllers.ReadHashtag)
	r.GET("/hashtags", controllers.ReadHashtags)
	r.POST("/hashtags", controllers.CreateHashtag)
	r.PUT("/hashtags/:id", controllers.UpdateHashtag)
	//hashtag project
	r.POST("/hashtagprojects", controllers.CreateProjectHashtags)
	//search
	r.GET("/search", controllers.SearchByUsername)
	r.Run(":5000")

}
