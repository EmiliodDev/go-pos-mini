package products

import (
	"context"

	repo "github.com/EmiliodDev/go-pos/internal/adapters/pgdb/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) CreateProduct(ctx context.Context, productParams repo.CreateProductParams) (repo.Product, error) {
	return repo.Product{}, nil
}
