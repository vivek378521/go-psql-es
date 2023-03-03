package controllers

import (
	"errors"
	"net/http"

	"example.com/go-psql-es/database"
	"example.com/go-psql-es/models"
	"github.com/gin-gonic/gin"
)

func CreateHashtag(c *gin.Context) {
	var hashtag *models.Hashtag
	err := c.ShouldBind(&hashtag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(hashtag)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a hashtag",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"hashtag": hashtag,
	})
	return
}

func ReadHashtag(c *gin.Context) {
	var hashtag models.Hashtag
	id := c.Param("id")
	res := database.DB.Find(&hashtag, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "hashtag not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"hashtag": hashtag,
	})
	return
}

func ReadHashtags(c *gin.Context) {
	var hashtags []models.Hashtag
	res := database.DB.Find(&hashtags)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("hashtags not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"hashtags": hashtags,
	})
	return
}

func UpdateHashtag(c *gin.Context) {
	var hashtag models.Hashtag
	id := c.Param("id")
	err := c.ShouldBind(&hashtag)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updatehashtag models.Hashtag
	res := database.DB.Model(&updatehashtag).Where("id = ?", id).Updates(hashtag)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "hashtag not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"hashtag": hashtag,
	})
	return
}
