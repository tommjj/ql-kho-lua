package handlers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type StorehouseHandler struct {
	scv ports.IStorehouseService
	acc ports.IAccessControlRepository
}

func NewStorehouseHandler(storehouseService ports.IStorehouseService, accessControl ports.IAccessControlRepository) *StorehouseHandler {
	return &StorehouseHandler{
		scv: storehouseService,
		acc: accessControl,
	}
}

type createStorehouseRequest struct {
	Name     string    `json:"name" binding:"required,min=3" example:"store 01"`
	Location []float64 `json:"location" binding:"required,location" example:"50.12,68.36"`
	Image    string    `json:"image" binding:"required,image_file" example:"2455.png"`
	Capacity int       `json:"capacity" binding:"required,min=1" example:"1200"`
}

// CreateStorehouse ql-kho-lua
//
//	@Summary		Create a new storehouse and get created user data
//	@Description	Create a new storehouse and get created storehouse data
//	@Tags			storehouses
//	@Accept			json
//	@Produce		json
//	@Param			request	body		createStorehouseRequest				true	"Create storehouse body"
//	@Success		200		{object}	response{data=storehouseResponse}	"Created storehouse data"
//	@Failure		400		{object}	errorResponse						"Validation error"
//	@Failure		401		{object}	errorResponse						"Unauthorized error"
//	@Failure		403		{object}	errorResponse						"Forbidden error"
//	@Failure		409		{object}	errorResponse						"Conflicting data error"
//	@Failure		500		{object}	errorResponse						"Internal server error"
//	@Router			/storehouses  [post]
//	@Security		JWTAuth
func (s *StorehouseHandler) CreateStorehouse(ctx *gin.Context) {
	var req createStorehouseRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	createdStore, err := s.scv.CreateStorehouse(ctx, &domain.Storehouse{
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

	res := newStorehouseResponse(createdStore)

	handleSuccess(ctx, res)
}

// GetStorehouseByID ql-kho-lua
//
//	@Summary		Get storehouse data
//	@Description	Get storehouse data by id
//	@Tags			storehouses
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int									true	"Storehouse id"
//	@Success		200	{object}	response{data=storehouseResponse}	"Created storehouse data"
//	@Failure		400	{object}	errorResponse						"Validation error"
//	@Failure		401	{object}	errorResponse						"Unauthorized error"
//	@Failure		403	{object}	errorResponse						"Forbidden error"
//	@Failure		404	{object}	errorResponse						"Data not found error"
//	@Failure		500	{object}	errorResponse						"Internal server error"
//	@Router			/storehouses/{id}  [get]
//	@Security		JWTAuth
func (s *StorehouseHandler) GetStorehouseByID(ctx *gin.Context) {
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

	store, err := s.scv.GetStorehouseByID(ctx, numID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newStorehouseResponse(store)
	handleSuccess(ctx, res)
}

type getListStorehouseRequest struct {
	Query string `form:"q" binding:"" example:"store 01"`
	Skip  int    `form:"skip" binding:"min=1" example:"1"`
	Limit int    `form:"limit" binding:"min=5" example:"5"`
}

// GetListStorehouses ql-kho-lua
//
//	@Summary		get storehouses
//	@Description	get storehouses
//	@Tags			storehouses
//	@Accept			json
//	@Produce		json
//	@Param			q		query		string										false	"Query"
//	@Param			skip	query		int											false	"Skip"	default(1)	minimum(1)
//	@Param			limit	query		int											false	"Limit"	default(5)	minimum(5)
//	@Success		200		{object}	responseWithPagination{data=[]userResponse}	"Storehouses data"
//	@Failure		400		{object}	errorResponse								"Validation error"
//	@Failure		401		{object}	errorResponse								"Unauthorized error"
//	@Failure		403		{object}	errorResponse								"Forbidden error"
//	@Failure		404		{object}	errorResponse								"Data not found error"
//	@Failure		500		{object}	errorResponse								"Internal server error"
//	@Router			/storehouses [get]
//	@Security		JWTAuth
func (s *StorehouseHandler) GetListStorehouses(ctx *gin.Context) {
	req := getListStorehouseRequest{
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
		count, err = s.scv.CountStorehouses(ctx, req.Query)
	} else {
		count, err = s.scv.CountAuthorizedStorehouses(ctx, token.ID, req.Query)
	}
	if err != nil {
		handleError(ctx, err)
		return
	}

	if checkPageOverflow(count, req.Limit, req.Skip) {
		handleError(ctx, domain.ErrDataNotFound)
		return
	}

	var stores []domain.Storehouse

	if isRoot {
		stores, err = s.scv.GetListStorehouses(ctx, req.Query, req.Limit, req.Skip)
	} else {
		stores, err = s.scv.GetAuthorizedStorehouses(ctx, token.ID, req.Query, req.Limit, req.Skip)
	}

	if err != nil {
		handleError(ctx, err)
		return
	}

	res := make([]storehouseResponse, 0, len(stores))

	for _, store := range stores {
		res = append(res, newStorehouseResponse(&store))
	}

	pagination := newPagination(count, len(stores), req.Limit, req.Skip)
	handleSuccessPagination(ctx, pagination, res)
}

type updateStorehouseRequest struct {
	Name     string    `json:"name" binding:"omitempty,min=3" example:"store 01"`
	Location []float64 `json:"location" binding:"omitempty,location" example:"51.12,68.36"`
	Image    string    `json:"image" binding:"omitempty,image_file" example:"2455.png"`
	Capacity int       `json:"capacity" binding:"omitempty,min=1" example:"1200"`
}

// UpdateStorehouse ql-kho-lua
//
//	@Summary		Update a storehouse and get created user data
//	@Description	Update a storehouse and get created storehouse data
//	@Tags			storehouses
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int									true	"Storehouse id"
//	@Param			request	body		updateStorehouseRequest				true	"Update storehouse body"
//	@Success		200		{object}	response{data=storehouseResponse}	"Updated storehouse data"
//	@Failure		400		{object}	errorResponse						"Validation error"
//	@Failure		401		{object}	errorResponse						"Unauthorized error"
//	@Failure		403		{object}	errorResponse						"Forbidden error"
//	@Failure		404		{object}	errorResponse						"Data not found error"
//	@Failure		409		{object}	errorResponse						"Conflicting data error"
//	@Failure		500		{object}	errorResponse						"Internal server error"
//	@Router			/storehouses/{id}  [patch]
//	@Security		JWTAuth
func (s *StorehouseHandler) UpdateStorehouse(ctx *gin.Context) {
	var req updateStorehouseRequest

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

	store, err := s.scv.UpdateStorehouse(ctx, &domain.Storehouse{
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

	res := newStorehouseResponse(store)
	handleSuccess(ctx, res)
}

// DeleteStorehouse ql-kho-lua
//
//	@Summary		Delete a storehouse
//	@Description	Delete a storehouse
//	@Tags			storehouses
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int				true	"Storehouse id"
//	@Success		200	{object}	response		"Updated storehouse data"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		401	{object}	errorResponse	"Unauthorized error"
//	@Failure		403	{object}	errorResponse	"Forbidden error"
//	@Failure		404	{object}	errorResponse	"Data not found error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/storehouses/{id}  [delete]
//	@Security		JWTAuth
func (s *StorehouseHandler) DeleteStorehouse(ctx *gin.Context) {
	id := ctx.Param("id")

	numID, err := strconv.Atoi(id)
	if err != nil {
		validationError(ctx, errors.New("id must be a number"))
		return
	}

	err = s.scv.DeleteStorehouse(ctx, numID)
	if err != nil {
		handleError(ctx, err)
		return
	}
	handleSuccess(ctx, nil)
}
