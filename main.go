package main

import (
	"fmt"

	"example.com/go-psql-es/controllers"
	"example.com/go-psql-es/database"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	r := gin.Default()
	r.GET("/users/:id", controllers.ReadUser)
	r.GET("/users", controllers.ReadUsers)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	//r.DELETE("/users/:id", controllers.DeleteBook)
	r.Run(":5000")
}
