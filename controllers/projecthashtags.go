package controllers

import (
	"net/http"

	"example.com/go-psql-es/database"
	"example.com/go-psql-es/models"
	"github.com/gin-gonic/gin"
)

func CreateProjectHashtags(c *gin.Context) {
	var hashtagproject *models.HashtagProject
	var project models.Project
	var hashtag models.Hashtag
	err := c.ShouldBind(&hashtagproject)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	hashtagFound := database.DB.Find(&hashtag, hashtagproject.HashtagId)
	if hashtagFound.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "hashtag not found",
		})
		return
	}
	projectFound := database.DB.Find(&project, hashtagproject.ProjectId)
	if projectFound.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "project not found",
		})
		return
	}
	res := database.DB.Create(hashtagproject)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a hashtagproject",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"hashtagproject": hashtagproject,
	})
	return
}
