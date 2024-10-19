package handlers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type WarehouseHandler struct {
	scv ports.IWarehouseService
	acc ports.IAccessControlRepository
}

func NewWarehouseHandler(warehouseService ports.IWarehouseService, accessControl ports.IAccessControlRepository) *WarehouseHandler {
	return &WarehouseHandler{
		scv: warehouseService,
		acc: accessControl,
	}
}

type createWarehouseRequest struct {
	Name     string    `json:"name" binding:"required,min=3,max=255" example:"store 01"`
	Location []float64 `json:"location" binding:"required,location" example:"50.12,68.36"`
	Image    string    `json:"image" binding:"required,image_file" example:"2455.png"`
	Capacity int       `json:"capacity" binding:"required,min=1" example:"1200"`
}

// CreateWarehouse ql-kho-lua
//
//	@Summary		Create a new warehouse and get created warehouse data
//	@Description	Create a new warehouse and get created warehouse data
//	@Tags			warehouses
//	@Accept			json
//	@Produce		json
//	@Param			request	body		createWarehouseRequest				true	"Create warehouse body"
//	@Success		200		{object}	response{data=warehouseResponse}	"Created warehouse data"
//	@Failure		400		{object}	errorResponse						"Validation error"
//	@Failure		401		{object}	errorResponse						"Unauthorized error"
//	@Failure		403		{object}	errorResponse						"Forbidden error"
//	@Failure		409		{object}	errorResponse						"Conflicting data error"
//	@Failure		500		{object}	errorResponse						"Internal server error"
//	@Router			/warehouses  [post]
//	@Security		JWTAuth
func (s *WarehouseHandler) CreateWarehouse(ctx *gin.Context) {
	var req createWarehouseRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	createdStore, err := s.scv.CreateWarehouse(ctx, &domain.Warehouse{
		Name:     req.Name,
		Location: fmt.Sprintf("%v, %v", req.Location[0], req.Location[1]),
		Capacity: req.Capacity,
		Image:    req.Image,
	})
	if err != nil {
		switch err {
		case domain.ErrFileIsNotExist:
			validationError(ctx, err)
		default:
			handleError(ctx, err)
		}
		return
	}

	res := newWarehouseResponse(createdStore)

	handleSuccess(ctx, res)
}

// GetWarehouseByID ql-kho-lua
//
//	@Summary		Get warehouse data
//	@Description	Get warehouse data by id
//	@Tags			warehouses
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int									true	"Warehouse id"
//	@Success		200	{object}	response{data=warehouseResponse}	"Warehouse data"
//	@Failure		400	{object}	errorResponse						"Validation error"
//	@Failure		401	{object}	errorResponse						"Unauthorized error"
//	@Failure		403	{object}	errorResponse						"Forbidden error"
//	@Failure		404	{object}	errorResponse						"Data not found error"
//	@Failure		500	{object}	errorResponse						"Internal server error"
//	@Router			/warehouses/{id}  [get]
//	@Security		JWTAuth
func (s *WarehouseHandler) GetWarehouseByID(ctx *gin.Context) {
	id := ctx.Param("id")

	numID, err := strconv.Atoi(id)
	if err != nil {
		validationError(ctx, errors.New("id must be a number"))
		return
	}

	token := getAuthPayload(ctx, authorizationPayloadKey)

	isRoot := token.Role == domain.Root

	if !isRoot {
		err := s.acc.HasAccess(ctx, numID, token.ID)
		if err != nil {
			handleError(ctx, err)
			return
		}
	}

	store, err := s.scv.GetWarehouseByID(ctx, numID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newWarehouseResponse(store)
	handleSuccess(ctx, res)
}

type getListWarehouseRequest struct {
	Query string `form:"q" binding:"" example:"store 01"`
	Skip  int    `form:"skip" binding:"min=1" example:"1"`
	Limit int    `form:"limit" binding:"min=5" example:"5"`
}

// GetListWarehouses ql-kho-lua
//
//	@Summary		get warehouses
//	@Description	get warehouses
//	@Tags			warehouses
//	@Accept			json
//	@Produce		json
//	@Param			q		query		string										false	"Query"
//	@Param			skip	query		int											false	"Skip"	default(1)	minimum(1)
//	@Param			limit	query		int											false	"Limit"	default(5)	minimum(5)
//	@Success		200		{object}	responseWithPagination{data=[]userResponse}	"Warehouses data"
//	@Failure		400		{object}	errorResponse								"Validation error"
//	@Failure		401		{object}	errorResponse								"Unauthorized error"
//	@Failure		403		{object}	errorResponse								"Forbidden error"
//	@Failure		404		{object}	errorResponse								"Data not found error"
//	@Failure		500		{object}	errorResponse								"Internal server error"
//	@Router			/warehouses [get]
//	@Security		JWTAuth
func (s *WarehouseHandler) GetListWarehouses(ctx *gin.Context) {
	req := getListWarehouseRequest{
		Limit: 5,
		Skip:  1,
	}
	err := ctx.BindQuery(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	token := getAuthPayload(ctx, authorizationPayloadKey)

	isRoot := token.Role == domain.Root

	var count int64

	if isRoot {
		count, err = s.scv.CountWarehouses(ctx, req.Query)
	} else {
		count, err = s.scv.CountAuthorizedWarehouses(ctx, token.ID, req.Query)
	}
	if err != nil {
		handleError(ctx, err)
		return
	}

	if checkPageOverflow(count, req.Limit, req.Skip) {
		handleError(ctx, domain.ErrDataNotFound)
		return
	}

	var stores []domain.Warehouse

	if isRoot {
		stores, err = s.scv.GetListWarehouses(ctx, req.Query, req.Limit, req.Skip)
	} else {
		stores, err = s.scv.GetAuthorizedWarehouses(ctx, token.ID, req.Query, req.Limit, req.Skip)
	}

	if err != nil {
		handleError(ctx, err)
		return
	}

	res := make([]warehouseResponse, 0, len(stores))

	for _, store := range stores {
		res = append(res, newWarehouseResponse(&store))
	}

	pagination := newPagination(count, len(stores), req.Limit, req.Skip)
	handleSuccessPagination(ctx, pagination, res)
}

// GetUsedCapacityByID ql-kho-lua
//
//	@Summary		Get used capacity
//	@Description	Get used capacity of warehouse by id
//	@Tags			warehouses
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int									true	"Warehouse id"
//	@Success		200	{object}	response{data=usedCapacityResponse}	"used capacity data"
//	@Failure		400	{object}	errorResponse						"Validation error"
//	@Failure		401	{object}	errorResponse						"Unauthorized error"
//	@Failure		403	{object}	errorResponse						"Forbidden error"
//	@Failure		404	{object}	errorResponse						"Data not found error"
//	@Failure		500	{object}	errorResponse						"Internal server error"
//	@Router			/warehouses/{id}/used_capacity  [get]
//	@Security		JWTAuth
func (s *WarehouseHandler) GetUsedCapacityByID(ctx *gin.Context) {
	id := ctx.Param("id")

	numID, err := strconv.Atoi(id)
	if err != nil {
		validationError(ctx, errors.New("id must be a number"))
		return
	}

	token := getAuthPayload(ctx, authorizationPayloadKey)

	isRoot := token.Role == domain.Root

	if !isRoot {
		err := s.acc.HasAccess(ctx, numID, token.ID)
		if err != nil {
			handleError(ctx, err)
			return
		}
	}

	usedCapacity, err := s.scv.GetUsedCapacityByID(ctx, numID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newUsedCapacityResponse(usedCapacity)
	handleSuccess(ctx, res)
}

type updateWarehouseRequest struct {
	Name     string    `json:"name" binding:"omitempty,min=3,max=255" example:"store 01"`
	Location []float64 `json:"location" binding:"omitempty,location" example:"51.12,68.36"`
	Image    string    `json:"image" binding:"omitempty,image_file" example:"2455.png"`
	Capacity int       `json:"capacity" binding:"omitempty,min=1" example:"1200"`
}

// UpdateWarehouse ql-kho-lua
//
//	@Summary		Update a warehouse and get created user data
//	@Description	Update a warehouse and get created warehouse data
//	@Tags			warehouses
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int									true	"Warehouse id"
//	@Param			request	body		updateWarehouseRequest				true	"Update warehouse body"
//	@Success		200		{object}	response{data=warehouseResponse}	"Updated warehouse data"
//	@Failure		400		{object}	errorResponse						"Validation error"
//	@Failure		401		{object}	errorResponse						"Unauthorized error"
//	@Failure		403		{object}	errorResponse						"Forbidden error"
//	@Failure		404		{object}	errorResponse						"Data not found error"
//	@Failure		409		{object}	errorResponse						"Conflicting data error"
//	@Failure		500		{object}	errorResponse						"Internal server error"
//	@Router			/warehouses/{id}  [patch]
//	@Security		JWTAuth
func (s *WarehouseHandler) UpdateWarehouse(ctx *gin.Context) {
	var req updateWarehouseRequest

	id := ctx.Param("id")

	numID, err := strconv.Atoi(id)
	if err != nil {
		validationError(ctx, errors.New("id must be a number"))
		return
	}

	err = ctx.BindJSON(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	if req.Capacity == 0 && len(req.Location) == 0 && req.Name == "" && req.Image == "" {
		handleError(ctx, domain.ErrNoUpdatedData)
		return
	}

	location := ""
	if len(req.Location) == 2 {
		location = fmt.Sprintf("%v, %v", req.Location[0], req.Location[1])
	}

	store, err := s.scv.UpdateWarehouse(ctx, &domain.Warehouse{
		ID:       numID,
		Name:     req.Name,
		Location: location,
		Capacity: req.Capacity,
		Image:    req.Image,
	})
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newWarehouseResponse(store)
	handleSuccess(ctx, res)
}

// DeleteWarehouse ql-kho-lua
//
//	@Summary		Delete a warehouse
//	@Description	Delete a warehouse
//	@Tags			warehouses
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int				true	"Warehouse id"
//	@Success		200	{object}	response		"deleted"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		401	{object}	errorResponse	"Unauthorized error"
//	@Failure		403	{object}	errorResponse	"Forbidden error"
//	@Failure		404	{object}	errorResponse	"Data not found error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/warehouses/{id}  [delete]
//	@Security		JWTAuth
func (s *WarehouseHandler) DeleteWarehouse(ctx *gin.Context) {
	id := ctx.Param("id")

	numID, err := strconv.Atoi(id)
	if err != nil {
		validationError(ctx, errors.New("id must be a number"))
		return
	}

	err = s.scv.DeleteWarehouse(ctx, numID)
	if err != nil {
		handleError(ctx, err)
		return
	}
	handleSuccess(ctx, nil)
}
