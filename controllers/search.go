package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"example.com/go-psql-es/database"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gin-gonic/gin"
)

type Document struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	ProjectName string   `json:"projectname"`
	Description string   `json:"description"`
	Slug        string   `json:"slug"`
	Hashtags    []string `json:"hashtags"`
}

func buildFuzzyQuery(searchText string) []byte {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"slug": map[string]interface{}{
								"query":     searchText,
								"fuzziness": "AUTO",
							},
						},
					},
					{
						"match": map[string]interface{}{
							"description": map[string]interface{}{
								"query":     searchText,
								"fuzziness": "AUTO",
							},
						},
					},
				},
			},
		},
	}

	q, err := json.Marshal(query)
	if err != nil {
		return nil
	}

	return q
}

func buildQuery(username string, hashtags []string) []byte {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []interface{}{},
			},
		},
	}

	if username != "" {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
			query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{}),
			map[string]interface{}{
				"match": map[string]interface{}{
					"username": username,
				},
			},
		)
	}

	if len(hashtags) > 0 && hashtags[0] != "" {
		for _, hashtag := range hashtags {
			query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
				query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{}),
				map[string]interface{}{
					"match": map[string]interface{}{
						"hashtags": hashtag,
					},
				},
			)
		}
	}
	fmt.Println(query)
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	return buf.Bytes()
}

func SearchByUsername(c *gin.Context) {
	var query []byte
	searchText := c.Query("searchText")
	userName := c.Query("username")
	hashtags := strings.Split(c.Query("hashtag"), ",")
	if searchText != "" {
		query = buildFuzzyQuery(searchText)
	}
	if userName != "" || hashtags[0] != "" {
		query = buildQuery(userName, hashtags)
	}

	req := esapi.SearchRequest{
		Index: []string{"myindex"},
		Body:  bytes.NewReader(query),
	}

	res, err := req.Do(context.Background(), database.EsClient)
	if err != nil {
		log.Fatalf("Error performing search: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error performing search: %s", res)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing search results: %s", err)
	}

	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	var documents []Document
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		doc := Document{
			ID:          source["id"].(string),
			Username:    source["username"].(string),
			ProjectName: source["projectname"].(string),
			Description: source["description"].(string),
			Slug:        source["slug"].(string),
			Hashtags:    toStringSlice(source["hashtags"]),
		}
		documents = append(documents, doc)
	}

	c.JSON(200, gin.H{
		"results": documents,
	})
}

func toStringSlice(interfaces interface{}) []string {
	result := make([]string, 0)
	for _, iface := range interfaces.([]interface{}) {
		result = append(result, iface.(string))
	}
	return result
}
