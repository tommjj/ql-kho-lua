package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type ImportInvoiceHandler struct {
	svc ports.IImportInvoicesService
}

func NewImportInvoiceHandler(svc ports.IImportInvoicesService) *ImportInvoiceHandler {
	return &ImportInvoiceHandler{
		svc: svc,
	}
}

type DetailInvoiceRequest struct {
	ID int `json:"id"`
}

func (i *ImportInvoiceHandler) CreateImInvoice(ctx *gin.Context) {

}
