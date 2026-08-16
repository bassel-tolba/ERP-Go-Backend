package service

import (
	"context"
	"fmt"

	"github.com/bassel-tolba/go-erp/internal/db"
)

type CompanyService struct {
	q *db.Queries
}

func NewCompanyService(q *db.Queries) *CompanyService {
	return &CompanyService{q: q}
}

func (s *CompanyService) CreateCompany(ctx context.Context, name string) (db.Company, error) {
	if name == "" {
		return db.Company{}, fmt.Errorf("company name is required")
	}
	return s.q.CreateCompany(ctx, name)
}
