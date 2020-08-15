package esclient

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v6"
	"log"
	"strings"
)

type EsClient struct {
	Client *elasticsearch.Client
	Config *elasticsearch.Config
}

func (e *EsClient) IntClient() error {
	var (
		r   map[string]interface{}
		err error
	)

	e.Client, err = elasticsearch.NewClient(*e.Config)

	// 1. Get cluster info
	res, err := e.Client.Info()
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// Check response status
	if res.IsError() {
		return err
	}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return err
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)

	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println("ElasticSearch client create succeed")

	log.Println(strings.Repeat("~", 37))
	return nil
}
