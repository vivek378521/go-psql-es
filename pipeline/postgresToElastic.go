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
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	ProjectName string   `json:"projectname"`
	Description string   `json:"description"`
	Slug        string   `json:"slug"`
	Hashtags    []string `json:"hashtags"`
}

func IngestToElastic(userProjectId int) {

	var userproject models.UserProject
	var project models.Project
	var user models.User
	var hashtagprojects1 []models.HashtagProject
	database.DB.Find(&userproject, userProjectId)
	database.DB.Find(&project, userproject.ProjectId)
	database.DB.Find(&user, userproject.UserId)
	database.DB.Where("project_id = ?", userproject.ProjectId).Find(&hashtagprojects1)
	var hashtagNames []string
	for _, hp := range hashtagprojects1 {
		var hashtag models.Hashtag
		database.DB.Find(&hashtag, hp.HashtagId)
		hashtagNames = append(hashtagNames, hashtag.Name)
	}

	// Print the collected hashtag names
	fmt.Println(hashtagNames)

	// Create an Elasticsearch data object
	esData := ElasticsearchData{
		ID:          fmt.Sprintf("%d", userproject.ID),
		Username:    user.Name,
		ProjectName: project.Name,
		Hashtags:    hashtagNames,
		Description: project.Description,
		Slug:        project.Slug,
	}
	fmt.Println(esData)
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
