package utils

import (
	"context"
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

func NewElasticClient() *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL(os.Getenv("ELASTIC_URL")))
	if err != nil {
		log.Fatalf("Error creating elastic client: %v", err)
	}
	return client
}

func PublishAccountEventES(client *elastic.Client, ctx context.Context, index string, id string, event string) error {
	_, err := client.Index().
		Index(index).
		Id(id).
		BodyJson(map[string]interface{}{"event": event}).
		Do(ctx)
	if err != nil {
		return err
	}
	return nil
}
