package controllers_test

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"example.com/go-psql-es/controllers"
	"example.com/go-psql-es/database"
	"example.com/go-psql-es/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestReadUser(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	database.DatabaseConnection()
	r := gin.Default()
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return
	}
	str := base64.URLEncoding.EncodeToString(b)
	user := &models.User{Name: str[:10]}
	fmt.Println("Idhr hu main")
	database.DB.Create(user)
	fmt.Println("Okay bas ")
	req, _ := http.NewRequest(http.MethodGet, "/users/"+strconv.FormatUint(uint64(user.ID), 10), nil)
	res := httptest.NewRecorder()

	// Test
	r.GET("/users/:id", controllers.ReadUser)
	r.ServeHTTP(res, req)

	// Assertions
	assert.Equal(t, http.StatusOK, res.Code)

	var response struct {
		User *models.User `json:"user"`
	}
	err = json.Unmarshal(res.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, user.Name, response.User.Name)
	//teardown
	database.DB.Where("id = ?", response.User.ID).Delete(&models.Project{})
}
