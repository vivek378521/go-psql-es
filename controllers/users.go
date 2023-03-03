package controllers

import (
	"errors"
	"net/http"

	"example.com/go-psql-es/database"
	"example.com/go-psql-es/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user *models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(user)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}

func ReadUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	res := database.DB.Find(&user, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}

func ReadUsers(c *gin.Context) {
	var users []models.User
	res := database.DB.Find(&users)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("users not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
	return
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	err := c.ShouldBind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updateUser models.User
	res := database.DB.Model(&updateUser).Where("id = ?", id).Updates(user)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}
