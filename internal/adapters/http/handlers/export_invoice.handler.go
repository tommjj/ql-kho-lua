package handlers

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type ExportInvoiceHandler struct {
	svc ports.IExportInvoiceService
	acc ports.IAccessControlService
}

func NewExportInvoiceHandler(svc ports.IExportInvoiceService, acc ports.IAccessControlService) *ExportInvoiceHandler {
	return &ExportInvoiceHandler{
		svc: svc,
		acc: acc,
	}
}

type detailExInvoiceRequest struct {
	RiceID   int     `json:"rice_id" binding:"required"`
	Price    float64 `json:"price" binding:"required,min=1"`
	Quantity int     `json:"quantity" binding:"required,min=1"`
}

type createExInvoiceRequest struct {
	WarehouseID int                      `json:"warehouse_id" binding:"required"`
	CustomerID  int                      `json:"customer_id" binding:"required"`
	Details     []detailExInvoiceRequest `json:"details" binding:"required,min=1,unique=RiceID"`
}

// CreateExInvoice ql-kho-lua
//
//	@Summary		Create a new export invoice and get created invoice data
//	@Description	Create a new export invoice and get created invoice data
//	@Tags			exportInvoices
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateExInvoiceRequest			true	"Create invoice body"
//	@Success		200		{object}	response{data=invoiceResponse}	"Created invoice data"
//	@Failure		400		{object}	errorResponse					"Validation error"
//	@Failure		401		{object}	errorResponse					"Unauthorized error"
//	@Failure		403		{object}	errorResponse					"Forbidden error"
//	@Failure		404		{object}	errorResponse					"Data not found error"
//	@Failure		409		{object}	errorResponse					"Conflicting data error"
//	@Failure		500		{object}	errorResponse					"Internal server error"
//	@Router			/export_invoices  [post]
//	@Security		JWTAuth
func (e *ExportInvoiceHandler) CreateExInvoice(ctx *gin.Context) {
	var req createExInvoiceRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	token := getAuthPayload(ctx, authorizationPayloadKey)

	isRootUser := token.Role == domain.Root
	if !isRootUser {
		err := e.acc.HasAccess(ctx, req.WarehouseID, token.ID)
		if err != nil {
			handleError(ctx, err)
			return
		}
	}

	createInvData := &domain.Invoice{
		WarehouseID: req.WarehouseID,
		CustomerID:  req.CustomerID,
		UserID:      token.ID,
		Details:     make([]domain.InvoiceItem, 0, len(req.Details)),
	}
	for _, v := range req.Details {
		createInvData.Details = append(createInvData.Details, domain.InvoiceItem{
			Price:    v.Price,
			Quantity: v.Quantity,
			RiceID:   v.RiceID,
		})
	}

	created, err := e.svc.CreateExInvoice(ctx, createInvData)
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newInvoiceResponse(created)
	handleSuccess(ctx, res)
}

// GetExInvoiceByID ql-kho-lua
//
//	@Summary		Get a export invoice by id
//	@Description	Get a export invoice by id
//	@Tags			exportInvoices
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int								true	"Invoice id"
//	@Success		200	{object}	response{data=invoiceResponse}	"Invoice data"
//	@Failure		400	{object}	errorResponse					"Validation error"
//	@Failure		401	{object}	errorResponse					"Unauthorized error"
//	@Failure		403	{object}	errorResponse					"Forbidden error"
//	@Failure		404	{object}	errorResponse					"Data not found error"
//	@Failure		500	{object}	errorResponse					"Internal server error"
//	@Router			/export_invoices/{id}  [get]
//	@Security		JWTAuth
func (e *ExportInvoiceHandler) GetExInvoiceByID(ctx *gin.Context) {
	id := ctx.Param("id")

	numID, err := strconv.Atoi(id)
	if err != nil {
		validationError(ctx, errors.New("id must be a number"))
		return
	}

	inv, err := e.svc.GetExInvoiceByID(ctx, numID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	token := getAuthPayload(ctx, authorizationPayloadKey)

	isRootUser := token.Role == domain.Root
	if !isRootUser {
		err := e.acc.HasAccess(ctx, inv.WarehouseID, token.ID)
		if err != nil {
			handleError(ctx, err)
			return
		}
	}

	res := newInvoiceResponse(inv)
	handleSuccess(ctx, res)
}

type getListExInvoiceRequest struct {
	WarehouseID int        `form:"warehouse_id" binding:"omitempty,min=0"`
	Start       *time.Time `form:"start" binding:"omitempty"`
	End         *time.Time `form:"end" binding:"omitempty"`
	Skip        int        `form:"skip" binding:"min=1" example:"1"`
	Limit       int        `form:"limit" binding:"min=5" example:"5"`
}

// GetListExInvoices ql-kho-lua
//
//	@Summary		Get a list export invoices
//	@Description	Get a list export invoices
//	@Tags			exportInvoices
//	@Accept			json
//	@Produce		json
//	@Param			warehouse_id	query		int												false	"Warehouse id"
//	@Param			skip			query		int												false	"Skip"	default(1)	minimum(1)
//	@Param			limit			query		int												false	"Limit"	default(5)	minimum(5)
//	@Param			start			query		string											false	"Start"	format(date-time)
//	@Param			end				query		string											false	"End"	format(date-time)
//	@Success		200				{object}	responseWithPagination{data=[]invoiceResponse}	"Invoice data"
//	@Failure		400				{object}	errorResponse									"Validation error"
//	@Failure		401				{object}	errorResponse									"Unauthorized error"
//	@Failure		403				{object}	errorResponse									"Forbidden error"
//	@Failure		404				{object}	errorResponse									"Data not found error"
//	@Failure		500				{object}	errorResponse									"Internal server error"
//	@Router			/export_invoices  [get]
//	@Security		JWTAuth
func (e *ExportInvoiceHandler) GetListExInvoices(ctx *gin.Context) {
	req := getListExInvoiceRequest{
		Skip:  1,
		Limit: 5,
	}
	err := ctx.BindQuery(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	token := getAuthPayload(ctx, authorizationPayloadKey)

	isRootUser := token.Role == domain.Root
	if !isRootUser {
		if req.WarehouseID == 0 {
			handleError(ctx, domain.ErrForbidden)
			return
		}
		err := e.acc.HasAccess(ctx, req.WarehouseID, token.ID)
		if err != nil {
			handleError(ctx, err)
			return
		}
	}

	count, err := e.svc.CountExInvoices(ctx, req.WarehouseID, req.Start, req.End)
	if err != nil {
		handleError(ctx, err)
		return
	}

	if checkPageOverflow(count, req.Limit, req.Skip) {
		handleError(ctx, domain.ErrDataNotFound)
		return
	}

	ivcs, err := e.svc.GetListExInvoices(ctx, req.WarehouseID, req.Start, req.End, req.Skip, req.Limit)
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := make([]invoiceResponse, 0, len(ivcs))
	for _, ivc := range ivcs {
		res = append(res, newInvoiceResponse(&ivc))
	}

	pagination := newPagination(count, len(ivcs), req.Limit, req.Skip)
	handleSuccessPagination(ctx, pagination, res)
}
