package pipeline

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"example.com/go-psql-es/database"
	"example.com/go-psql-es/models"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type ElasticsearchData struct {
	ID          string
	UserName    string
	ProjectName string
	Hashtags    []string
	Description string
	Slug        string
}

func IngestToElastic(userProjectId int) {

	var userproject models.UserProject
	var project models.Project
	var user models.User
	var hashtagprojects1 []models.HashtagProject
	database.DB.Find(&userproject, userProjectId)
	fmt.Println("*************")
	fmt.Println(userproject)
	database.DB.Find(&project, userproject.ProjectId)
	database.DB.Find(&user, userproject.UserId)
	fmt.Println(user)
	fmt.Println(project)
	database.DB.Where("project_id = ?", project.ID).Find(&hashtagprojects1)
	fmt.Println(hashtagprojects1)
	fmt.Println(len(hashtagprojects1))

	var hashtagNames []string
	for _, hp := range hashtagprojects1 {
		var hashtag models.Hashtag
		database.DB.Find(&hashtag, hp.ID)
		hashtagNames = append(hashtagNames, hashtag.Name)
	}

	// Print the collected hashtag names
	fmt.Println(hashtagNames)

	// Create an Elasticsearch data object
	esData := ElasticsearchData{
		ID:          fmt.Sprintf("%d", userproject.ID),
		UserName:    user.Name,
		ProjectName: project.Name,
		Hashtags:    hashtagNames,
		Description: project.Description,
		Slug:        project.Slug,
	}

	// Convert the data object to JSON
	jsonData, err := json.Marshal(esData)
	if err != nil {
		log.Fatalf("Error converting data to JSON: %s", err)
	}
	req := esapi.IndexRequest{
		Index:      "myindex",
		DocumentID: fmt.Sprintf("%d", userproject.ID),
		Body:       bytes.NewReader(jsonData),
	}

	// Send the request to Elasticsearch
	res, err := req.Do(context.Background(), database.EsClient)
	if err != nil {
		log.Fatalf("Error sending request to Elasticsearch: %s", err)
	}
	defer res.Body.Close()

	// Check the response status
	if res.IsError() {
		log.Fatalf("Error response received from Elasticsearch: %s", res.Status())
	}

	log.Printf("Data indexed successfully: %s", res.Status())

}
