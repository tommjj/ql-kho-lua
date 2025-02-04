package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type CustomerHandler struct {
	svc ports.ICustomerService
}

func NewCustomerHandler(customerService ports.ICustomerService) *CustomerHandler {
	return &CustomerHandler{
		svc: customerService,
	}
}

type createCustomerRequest struct {
	Name    string `json:"name" binding:"required,min=3,max=255" example:"Sentenced"`
	Email   string `json:"email" binding:"required,email" example:"example@exp.com"`
	Phone   string `json:"phone" binding:"required,e164" example:"+84123456789"`
	Address string `json:"address" binding:"required,min=1,max=255" example:"abc, xyz"`
}

// CreateCustomer ql-kho-lua
//
//	@Summary		Create a new customer
//	@Description	Create a new customer and get created customer data
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			request	body		createCustomerRequest			true	"Create customer body"
//	@Success		200		{object}	response{data=customerResponse}	"Created customer data"
//	@Failure		400		{object}	errorResponse					"Validation error"
//	@Failure		401		{object}	errorResponse					"Unauthorized error"
//	@Failure		403		{object}	errorResponse					"Forbidden error"
//	@Failure		409		{object}	errorResponse					"Conflicting data error"
//	@Failure		500		{object}	errorResponse					"Internal server error"
//	@Router			/customers  [post]
//	@Security		JWTAuth
func (c *CustomerHandler) CreateCustomer(ctx *gin.Context) {
	var req createCustomerRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	customer, err := c.svc.CreateCustomer(ctx, &domain.Customer{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	})
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newCustomerResponse(customer)
	handleSuccess(ctx, res)
}

// GetCustomerByID ql-kho-lua
//
//	@Summary		Get a customer
//	@Description	Get a customer by customer id
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int								true	"Customer id"
//	@Success		200	{object}	response{data=customerResponse}	"Customer data"
//	@Failure		400	{object}	errorResponse					"Validation error"
//	@Failure		401	{object}	errorResponse					"Unauthorized error"
//	@Failure		403	{object}	errorResponse					"Forbidden error"
//	@Failure		404	{object}	errorResponse					"Data not found error"
//	@Failure		500	{object}	errorResponse					"Internal server error"
//	@Router			/customers/{id}  [get]
//	@Security		JWTAuth
func (c *CustomerHandler) GetCustomerByID(ctx *gin.Context) {
	id := ctx.Param("id")

	numID, err := strconv.Atoi(id)
	if err != nil {
		validationError(ctx, errors.New("id must be a number"))
		return
	}

	customer, err := c.svc.GetCustomerByID(ctx, numID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newCustomerResponse(customer)
	handleSuccess(ctx, res)
}

type getListCustomerRequest struct {
	Query string `form:"q" binding:"" example:"teo"`
	Skip  int    `form:"skip" binding:"min=1" example:"1"`
	Limit int    `form:"limit" binding:"min=5" example:"5"`
}

// GetListCustomers ql-kho-lua
//
//	@Summary		get customers
//	@Description	get customers with pagination
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			q		query		string											false	"Query"
//	@Param			skip	query		int												false	"Skip"	default(1)	minimum(1)
//	@Param			limit	query		int												false	"Limit"	default(5)	minimum(5)
//	@Success		200		{object}	responseWithPagination{data=[]customerResponse}	"Customers data"
//	@Failure		400		{object}	errorResponse									"Validation error"
//	@Failure		401		{object}	errorResponse									"Unauthorized error"
//	@Failure		403		{object}	errorResponse									"Forbidden error"
//	@Failure		404		{object}	errorResponse									"Data not found error"
//	@Failure		500		{object}	errorResponse									"Internal server error"
//	@Router			/customers [get]
//	@Security		JWTAuth
func (c *CustomerHandler) GetListCustomers(ctx *gin.Context) {
	req := getListCustomerRequest{
		Skip:  1,
		Limit: 5,
	}
	err := ctx.BindQuery(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	count, err := c.svc.CountCustomers(ctx, req.Query)
	if err != nil {
		handleError(ctx, err)
		return
	}

	if checkPageOverflow(count, req.Limit, req.Skip) {
		handleError(ctx, domain.ErrDataNotFound)
		return
	}

	customers, err := c.svc.GetListCustomers(ctx, req.Query, req.Limit, req.Skip)
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := make([]customerResponse, 0, len(customers))
	for _, customer := range customers {
		res = append(res, newCustomerResponse(&customer))
	}

	pagination := newPagination(count, len(customers), req.Limit, req.Skip)
	handleSuccessPagination(ctx, pagination, res)
}

type updateCustomerRequest struct {
	Name    string `json:"name" binding:"omitempty,min=3,max=255" example:"Sentenced"`
	Email   string `json:"email" binding:"omitempty,email" example:"example@exp.com"`
	Phone   string `json:"phone" binding:"omitempty,e164" example:"+84123456789"`
	Address string `json:"address" binding:"omitempty,min=1,max=255" example:"abc, xyz"`
}

// UpdateCustomer ql-kho-lua
//
//	@Summary		update customer
//	@Description	update customer and get updated customer data
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int								true	"Customer id"
//	@Param			request	body		updateCustomerRequest			true	"Update customer body"
//	@Success		200		{object}	response{data=customerResponse}	"Updated customer data"
//	@Failure		400		{object}	errorResponse					"Validation error"
//	@Failure		401		{object}	errorResponse					"Unauthorized error"
//	@Failure		403		{object}	errorResponse					"Forbidden error"
//	@Failure		404		{object}	errorResponse					"Data not found error"
//	@Failure		500		{object}	errorResponse					"Internal server error"
//	@Router			/customers/{id} [patch]
//	@Security		JWTAuth
func (c *CustomerHandler) UpdateCustomer(ctx *gin.Context) {
	id := ctx.Param("id")

	numID, err := strconv.Atoi(id)
	if err != nil {
		validationError(ctx, errors.New("id must be a number"))
		return
	}

	var req updateCustomerRequest

	err = ctx.BindJSON(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	customer, err := c.svc.UpdateCustomer(ctx, &domain.Customer{
		ID:      numID,
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	})
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newCustomerResponse(customer)
	handleSuccess(ctx, res)
}

// DeleteCustomer ql-kho-lua
//
//	@Summary		delete customer
//	@Description	delete customer by id
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int				true	"Customer id"
//	@Success		200	{object}	response		"Deleted data"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		401	{object}	errorResponse	"Unauthorized error"
//	@Failure		403	{object}	errorResponse	"Forbidden error"
//	@Failure		404	{object}	errorResponse	"Data not found error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/customers/{id} [delete]
//	@Security		JWTAuth
func (c *CustomerHandler) DeleteCustomer(ctx *gin.Context) {
	id := ctx.Param("id")

	numID, err := strconv.Atoi(id)
	if err != nil {
		validationError(ctx, errors.New("id must be a number"))
		return
	}

	err = c.svc.DeleteCustomer(ctx, numID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, nil)
}
