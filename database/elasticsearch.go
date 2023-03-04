package database

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

var EsClient *elasticsearch.Client

func ElasticSearchConnection() {
	cfg := elasticsearch.Config{
		CloudID: "go-psql-es:dXMtY2VudHJhbDEuZ2NwLmNsb3VkLmVzLmlvJDI2ODI0MWIyZDNlYTRhMDJiYTk1NTI4ZmI5YTVlNmM5JGYzOWY5ZmUyYTE4YTQzYmNhZTJkMDQyYTc0ZmM1ZDcx",
		APIKey:  "UTVva3E0WUJrOTF5Mm5jY0FZLS06WEtqTG9GeHBUaE92cFlIY2RYVFJuUQ==",
	}
	EsClient, err = elasticsearch.NewClient(cfg)
	gucci, err := EsClient.Info()
	fmt.Println(gucci)
	res, err := EsClient.Ping()
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.IsError() {
		panic(fmt.Errorf("Error: %s", res.String()))
	}
	fmt.Println("ElasticSearch connection successful...")
}
