package handlers

import "github.com/tommjj/ql-kho-lua/internal/core/ports"

type CustomerHandler struct {
	svc ports.ICustomerService
}

func NewCustomerHandler(customerService ports.ICustomerService) *CustomerHandler {
	return &CustomerHandler{
		svc: customerService,
	}
}
