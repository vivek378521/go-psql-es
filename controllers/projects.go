package controllers

import (
	"errors"
	"net/http"

	"example.com/go-psql-es/database"
	"example.com/go-psql-es/models"
	"github.com/gin-gonic/gin"
)

type UpdateProjectInput struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

func CreateProject(c *gin.Context) {
	var project *models.Project
	err := c.ShouldBind(&project)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(project)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a project",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"project": project,
	})
	return
}

func ReadProject(c *gin.Context) {
	var project models.Project
	id := c.Param("id")
	res := database.DB.Find(&project, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "project not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"project": project,
	})
	return
}

func ReadProjects(c *gin.Context) {
	var projects []models.Project
	res := database.DB.Find(&projects)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("projects not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
	return
}

func UpdateProject(c *gin.Context) {
	var project models.Project
	id := c.Param("id")
	err := c.ShouldBind(&project)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updateproject models.Project
	res := database.DB.Model(&updateproject).Where("id = ?", id).Updates(project)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "project not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"project": project,
	})
	return
}

func PatchProject(c *gin.Context) {
	var project models.Project
	if err := database.DB.Where("id = ?", c.Param("id")).First(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateProject := models.Project{Name: input.Name, Slug: input.Slug, Description: input.Description}

	database.DB.Model(&project).Updates(updateProject)

	c.JSON(http.StatusOK, gin.H{"project": project})
}
