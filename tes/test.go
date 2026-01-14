package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {

	type Product struct {
		ID          string  `json:"id"`
		ProductName string  `json:"product_name"`
		DrugGeneric string  `json:"drug_generic"`
		Company     string  `json:"company"`
		Score       float64 `json:"score"`
	}

	type SearchResponse struct {
		Results []Product `json:"results"`
	}
	var raw struct {
		Hits struct {
			Hits []struct {
				Index  string         `json:"_index"`
				ID     string         `json:"_id"`
				Score  float64        `json:"_score"`
				Source map[string]any `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	strBody := strings.NewReader(`{"took":10,"timed_out":false,"_shards":{"total":1,"successful":1,"skipped":0,"failed":0},"hits":{"total":{"value":1,"relation":"eq"},"max_score":10.284903,"hits":[{"_index":"products","_id":"10","_score":10.284903,"_source":{  "id": "10",  "product_name": "carbidopa, levodopa and entacapone",  "drug_generic": "carbidopa, levodopa and entacapone",  "company": "Wockhardt USA LLC."}
}]}}}`)

	log.Print(strBody)
	if err := json.NewDecoder(strBody).Decode(&raw); err != nil {
		log.Println("decode error:", err)
		return
	}

	out := SearchResponse{Results: make([]Product, 0, len(raw.Hits.Hits))}
	for _, h := range raw.Hits.Hits {
		p := Product{
			ID:          toString(h.Source["id"]),
			ProductName: toString(h.Source["product_name"]),
			DrugGeneric: toString(h.Source["drug_generic"]),
			Company:     toString(h.Source["company"]),
			Score:       h.Score,
		}
		out.Results = append(out.Results, p)
	}

	fmt.Println("Search Results:", out)
}

func toString(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	b, _ := json.Marshal(v)
	return string(b)
}
