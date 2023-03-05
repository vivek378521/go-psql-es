package database

import (
	"fmt"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

var EsClient *elasticsearch.Client

func ElasticSearchConnection() {
	cfg := elasticsearch.Config{
		CloudID: os.Getenv("ES_CLOUD_ID"),
		APIKey:  os.Getenv("ES_CLOUD_API_KEY"),
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
