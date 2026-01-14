package service

import (
	"context"

	"github.com/baihakhi/product-search/internal/model"
	"github.com/baihakhi/product-search/internal/repository"
)

type (
	ProductService interface {
		SearchProducts(ctx context.Context, size int, q, idx string) ([]model.Product, error)
	}
	productService struct {
		repo repository.ProductRepo
	}
)

func NewProductService(repo repository.ProductRepo) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) SearchProducts(ctx context.Context, size int, q, idx string) ([]model.Product, error) {
	return s.repo.SearchProducts(ctx, size, q, idx)
}
