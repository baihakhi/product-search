package repository

import (
	"context"
	"encoding/json"

	"github.com/baihakhi/product-search/internal/config"
	"github.com/baihakhi/product-search/internal/model"
)

type ProductRepo interface {
	SearchProducts(ctx context.Context, size int, q, idx string) ([]model.Product, error)
}

type productRepo struct {
	esClient config.ESClient
}

func NewProductRepo(esClient config.ESClient) ProductRepo {
	return &productRepo{
		esClient: esClient,
	}
}

func (r *productRepo) SearchProducts(ctx context.Context, size int, q, idx string) ([]model.Product, error) {

	data, err := json.Marshal(map[string]any{
		"size": size,
		"query": map[string]any{
			"multi_match": map[string]any{
				"query":                q,
				"fields":               []string{"product_name^3", "drug_generic^2", "company"},
				"operator":             "or",
				"fuzziness":            "AUTO",
				"prefix_length":        1,
				"minimum_should_match": "1",
			},
		},
	})
	if err != nil {
		return nil, err
	}

	res, err := r.esClient.Search(ctx, data, idx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	rawResponse := model.RawResponseProduct{}
	if err := json.NewDecoder(res.Body).Decode(&rawResponse); err != nil {
		return nil, err
	}

	return model.MapToProducts(rawResponse), nil
}
