package elasticsearch

import (
	"context"
	"fmt"
	"github.com/Emanuel9/bookstore_items-api/logger"
	"github.com/olivere/elastic"
	"time"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10 * time.Second),
		elastic.SetSniff(false),
		
	)

	if err != nil {
		panic(err)
	}
	Client.setClient(client)

	// Create the index if it does not exist
}

func (c *esClient) Index(index string, doc interface{})(*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		BodyJson(doc).
		Do(ctx)

	if err != nil {
		logger.Error(
			fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}

	return result, nil
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}