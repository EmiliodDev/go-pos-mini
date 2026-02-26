package products

import (
	"context"

	repo "github.com/EmiliodDev/go-pos/internal/adapters/pgdb/sqlc"
)

// var errProductNotFound = errors.New("product not found")

type Service interface {
	listProducts(ctx context.Context) ([]repo.Product, error)
	createProduct(ctx context.Context, args repo.CreateProductParams) (repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) listProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) createProduct(ctx context.Context, args repo.CreateProductParams) (repo.Product, error) {
	product, err := s.repo.CreateProduct(ctx, args)
	if err != nil {
		return repo.Product{}, err
	}

	return product, nil
}
