package config

import (
	"context"
	"os"
	"strings"

	esv9 "github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"
)

var (
	ESIndexProducts string = os.Getenv("ES_INDEX")
)

type (
	ESClient interface {
		Search(ctx context.Context, body []byte, index string) (*esapi.Response, error)
	}
	esClient struct {
		Client *esv9.Client
	}
	ESConfig struct {
		CloudID string
		APIKey  string
	}
)

func NewESClient(client *esv9.Client) *esClient {
	return &esClient{
		Client: client,
	}
}

func LoadESConfig() *ESConfig {
	return &ESConfig{
		CloudID: os.Getenv("ES_CLOUD_ID"),
		APIKey:  os.Getenv("ES_API_KEY"),
	}
}

func InitESClient(cfg *ESConfig) (*esv9.Client, error) {
	esCfg := esv9.Config{
		CloudID: cfg.CloudID,
		APIKey:  cfg.APIKey,
	}
	client, err := esv9.NewClient(esCfg)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (e *esClient) Search(ctx context.Context, body []byte, index string) (*esapi.Response, error) {
	res, err := e.Client.Search(
		e.Client.Search.WithContext(ctx),
		e.Client.Search.WithIndex(index),
		e.Client.Search.WithBody(strings.NewReader(string(body))),
	)
	return res, err
}
