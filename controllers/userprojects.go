package controllers

import (
	"errors"
	"net/http"

	"example.com/go-psql-es/database"
	"example.com/go-psql-es/models"
	"example.com/go-psql-es/pipeline"
	"github.com/gin-gonic/gin"
)

type UpdateUserProjectInput struct {
	UserId    uint `json:"userId"`
	ProjectId uint `json:"projectId"`
}

func CreateUserProject(c *gin.Context) {
	var userproject *models.UserProject
	var user models.User
	var project models.Project
	err := c.ShouldBind(&userproject)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	userFound := database.DB.Find(&user, userproject.UserId)
	if userFound.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}
	projectFound := database.DB.Find(&project, userproject.ProjectId)
	if projectFound.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "project not found",
		})
		return
	}
	res := database.DB.Create(userproject)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a userproject",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userproject": userproject,
	})
	return
}

func ReadUserProject(c *gin.Context) {
	var userproject models.UserProject
	id := c.Param("id")
	res := database.DB.Find(&userproject, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "userproject not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userproject": userproject,
	})
	return
}

func ReadUserProjects(c *gin.Context) {
	var userprojects []models.UserProject
	res := database.DB.Find(&userprojects)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("userprojects not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userprojects": userprojects,
	})
	return
}

func UpdateUserProject(c *gin.Context) {
	var userproject models.UserProject
	var user models.User
	var project models.Project
	id := c.Param("id")
	err := c.ShouldBind(&userproject)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	userFound := database.DB.Find(&user, userproject.UserId)
	if userFound.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}
	projectFound := database.DB.Find(&project, userproject.ProjectId)
	if projectFound.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "project not found",
		})
		return
	}
	var updateuserproject models.UserProject
	res := database.DB.Model(&updateuserproject).Where("id = ?", id).Updates(userproject)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userproject not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userproject": userproject,
		"user":        userFound,
		"project":     projectFound,
	})
	return
}

func PatchUserProject(c *gin.Context) {
	var userproject models.UserProject
	var user models.User
	var project models.Project
	if err := database.DB.Where("id = ?", c.Param("id")).First(&userproject).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	userFound := database.DB.Find(&user, userproject.UserId)
	if userFound.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}
	projectFound := database.DB.Find(&project, userproject.ProjectId)
	if projectFound.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "project not found",
		})
		return
	}

	// Validate input
	var input UpdateUserProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateuserproject := models.UserProject{UserId: input.UserId, ProjectId: input.ProjectId}

	database.DB.Model(&userproject).Updates(updateuserproject)

	c.JSON(http.StatusOK, gin.H{"userproject": userproject, "user": userFound, "project": projectFound})
}

func Ingest(c *gin.Context) {
	pipeline.IngestToElastic(1)
}
