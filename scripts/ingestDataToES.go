package scripts

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"example.com/go-psql-es/database"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Document struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	ProjectName string   `json:"projectname"`
	Description string   `json:"description"`
	Slug        string   `json:"slug"`
	Hashtags    []string `json:"hashtags"`
}

func IngestBulkData() {

	// Create multiple JSON entries
	documents := []Document{
		{
			ID:          "doc4",
			Username:    "Vivek Khatri",
			ProjectName: "Trigger",
			Description: "This is a project that does fuzzy search",
			Slug:        "creating-oh-lol",
			Hashtags:    []string{"tag1", "tag2"},
		},
		{
			ID:          "doc5",
			Username:    "Vivek Khatri",
			ProjectName: "Dragon",
			Description: "This is a project that does fuzzy buzzy search",
			Slug:        "creating-oh-lol-1",
			Hashtags:    []string{"tag3", "tag4"},
		},
		{
			ID:          "doc6",
			Username:    "Zeus Sinha",
			ProjectName: "Chant",
			Description: "This is a project that does fuzzy buzzy guzzy search",
			Slug:        "creating-oh-lol-2",
			Hashtags:    []string{"tag1", "tag5"},
		},
	}

	// Create a bulk request object
	var buf bytes.Buffer

	for _, doc := range documents {
		// Serialize the document to JSON
		docJSON, err := json.Marshal(doc)
		if err != nil {
			log.Fatalf("Error serializing document %s: %s", doc.ID, err)
		}

		// Add the index operation to the bulk request object
		meta := []byte(fmt.Sprintf(`{ "index" : { "_index" : "projects", "_id" : "%s" } }%s`, doc.ID, "\n"))
		buf.Grow(len(meta) + len(docJSON))
		buf.Write(meta)
		buf.Write(docJSON)
		buf.WriteString("\n")
	}

	// Build the request object
	req := esapi.BulkRequest{
		Body: bytes.NewReader(buf.Bytes()),
	}

	// Execute the request
	res, err := req.Do(context.Background(), database.EsClient)
	if err != nil {
		log.Fatalf("Error indexing documents: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error indexing documents: %s", res.String())
	}

	fmt.Printf("Bulk ingestion successful with status code: %d\n", res.StatusCode)
}
