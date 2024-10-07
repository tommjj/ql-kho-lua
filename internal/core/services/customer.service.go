package services

import "github.com/tommjj/ql-kho-lua/internal/core/ports"

type customerService struct {
	repo ports.ICustomerRepository
}

func NewCustomerService(repo ports.ICustomerRepository) *customerService {
	return &customerService{
		repo: repo,
	}
}
