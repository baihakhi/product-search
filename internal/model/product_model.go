package model

import "github.com/baihakhi/product-search/internal/util"

type Product struct {
	ID          string  `json:"id"`
	ProductName string  `json:"product_name"`
	DrugGeneric string  `json:"drug_generic"`
	Company     string  `json:"company"`
	Score       float64 `json:"score"`
}

// ResponseProduct is the response structure for the API
type ResponseProduct struct {
	Results []Product `json:"results"`
}

// RawResponseProduct is used to unmarshal the raw Elasticsearch response
type RawResponseProduct struct {
	Hits struct {
		Hits []struct {
			Index  string         `json:"_index"`
			ID     string         `json:"_id"`
			Score  float64        `json:"_score"`
			Source map[string]any `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func MapToProducts(raw RawResponseProduct) []Product {
	products := make([]Product, 0, len(raw.Hits.Hits))
	for _, h := range raw.Hits.Hits {
		p := Product{
			ID:          util.ToString(h.Source["id"]),
			ProductName: util.ToString(h.Source["product_name"]),
			DrugGeneric: util.ToString(h.Source["drug_generic"]),
			Company:     util.ToString(h.Source["company"]),
			Score:       h.Score,
		}
		products = append(products, p)
	}
	return products
}
