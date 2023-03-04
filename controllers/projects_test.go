package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/go-psql-es/controllers"
	"example.com/go-psql-es/database"
	"example.com/go-psql-es/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateProject(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	database.DatabaseConnection()
	r := gin.Default()
	project := &models.Project{Name: "Test Project"}
	reqBody, _ := json.Marshal(project)
	req, _ := http.NewRequest(http.MethodPost, "/projects", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	// Test
	r.POST("/projects", controllers.CreateProject)
	r.ServeHTTP(res, req)

	// Assertions
	assert.Equal(t, http.StatusOK, res.Code)

	var response struct {
		Project *models.Project `json:"project"`
	}
	err := json.Unmarshal(res.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, project.Name, response.Project.Name)

	// Teardown
	database.DB.Where("id = ?", response.Project.ID).Delete(&models.Project{})
}
