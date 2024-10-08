package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type RiceHandler struct {
	svc ports.IRiceService
}

func NewRiceHandler(riceService ports.IRiceService) *RiceHandler {
	return &RiceHandler{
		svc: riceService,
	}
}

type createRiceRequest struct {
	Name string `json:"name" binding:"required"`
}

// CreateRice ql-kho-lua
//
//	@Summary		Create a new rice and get created rice data
//	@Description	Create a new rice and get created rice data
//	@Tags			rice
//	@Accept			json
//	@Produce		json
//	@Param			request	body		createRiceRequest			true	"Create rice body"
//	@Success		200		{object}	response{data=riceResponse}	"Created rice data"
//	@Failure		400		{object}	errorResponse				"Validation error"
//	@Failure		401		{object}	errorResponse				"Unauthorized error"
//	@Failure		403		{object}	errorResponse				"Forbidden error"
//	@Failure		409		{object}	errorResponse				"Conflicting data error"
//	@Failure		500		{object}	errorResponse				"Internal server error"
//	@Router			/rice  [post]
//	@Security		JWTAuth
func (r *RiceHandler) CreateRice(ctx *gin.Context) {
	var req createRiceRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	rice, err := r.svc.CreateRice(ctx, &domain.Rice{
		Name: req.Name,
	})
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newRiceResponse(rice)
	handleSuccess(ctx, res)
}

// GetRiceByID ql-kho-lua
//
//	@Summary		Get a rice
//	@Description	Get a rice by user id
//	@Tags			rice
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int							true	"Rice id"
//	@Success		200	{object}	response{data=riceResponse}	"Rice data"
//	@Failure		400	{object}	errorResponse				"Validation error"
//	@Failure		401	{object}	errorResponse				"Unauthorized error"
//	@Failure		404	{object}	errorResponse				"Data not found error"
//	@Failure		500	{object}	errorResponse				"Internal server error"
//	@Router			/rice/{id}  [get]
//	@Security		JWTAuth
func (r *RiceHandler) GetRiceByID(ctx *gin.Context) {
	id := ctx.Param("id")

	numID, err := strconv.Atoi(id)
	if err != nil {
		validationError(ctx, errors.New("id must be a number"))
		return
	}

	rice, err := r.svc.GetRiceByID(ctx, numID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newRiceResponse(rice)
	handleSuccess(ctx, res)
}

type getListRiceRequest struct {
	Query string `form:"q" binding:"" example:"teo"`
	Skip  int    `form:"skip" binding:"min=1" example:"1"`
	Limit int    `form:"limit" binding:"min=5" example:"5"`
}

// GetListRice ql-kho-lua
//
//	@Summary		get list rice
//	@Description	get list rice
//	@Tags			rice
//	@Accept			json
//	@Produce		json
//	@Param			q		query		string										false	"Query"
//	@Param			skip	query		int											false	"Skip"	default(1)	minimum(1)
//	@Param			limit	query		int											false	"Limit"	default(5)	minimum(5)
//	@Success		200		{object}	responseWithPagination{data=[]riceResponse}	"Rice data"
//	@Failure		400		{object}	errorResponse								"Validation error"
//	@Failure		401		{object}	errorResponse								"Unauthorized error"
//	@Failure		404		{object}	errorResponse								"Data not found error"
//	@Failure		500		{object}	errorResponse								"Internal server error"
//	@Router			/rice [get]
//	@Security		JWTAuth
func (r *RiceHandler) GetListRice(ctx *gin.Context) {
	req := getListRiceRequest{
		Limit: 5,
		Skip:  1,
	}
	err := ctx.BindQuery(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	count, err := r.svc.CountRice(ctx, req.Query)
	if err != nil {
		handleError(ctx, err)
		return
	}

	if checkPageOverflow(count, req.Limit, req.Skip) {
		handleError(ctx, domain.ErrDataNotFound)
		return
	}

	rice, err := r.svc.GetListRice(ctx, req.Query, req.Limit, req.Skip)
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := make([]riceResponse, 0, len(rice))
	for _, item := range rice {
		res = append(res, riceResponse(item))
	}

	pagination := newPagination(count, len(rice), req.Limit, req.Skip)
	handleSuccessPagination(ctx, pagination, res)
}

type updateRiceRequest struct {
	Name string `json:"name" binding:"required"`
}

// UpdateRice ql-kho-lua
//
//	@Summary		update rice
//	@Description	update rice
//	@Tags			rice
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Rice id"
//	@Param			request	body		updateRiceRequest			true	"Update rice body"
//	@Success		200		{object}	response{data=riceResponse}	"Updated rice data"
//	@Failure		400		{object}	errorResponse				"Validation error"
//	@Failure		401		{object}	errorResponse				"Unauthorized error"
//	@Failure		403		{object}	errorResponse				"Forbidden error"
//	@Failure		404		{object}	errorResponse				"Data not found error"
//	@Failure		409		{object}	errorResponse				"Conflicting data error"
//	@Failure		500		{object}	errorResponse				"Internal server error"
//	@Router			/rice/{id} [patch]
//	@Security		JWTAuth
func (r *RiceHandler) UpdateRice(ctx *gin.Context) {
	var req updateRiceRequest
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

	rice, err := r.svc.UpdateRice(ctx, &domain.Rice{
		ID:   numID,
		Name: req.Name,
	})
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newRiceResponse(rice)
	handleSuccess(ctx, res)
}

// DeleteRice ql-kho-lua
//
//	@Summary		delete rice
//	@Description	delete rice
//	@Tags			rice
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int				true	"Rice id"
//	@Success		200	{object}	response		"Deleted data"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		401	{object}	errorResponse	"Unauthorized error"
//	@Failure		403	{object}	errorResponse	"Forbidden error"
//	@Failure		404	{object}	errorResponse	"Data not found error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/rice/{id} [delete]
//	@Security		JWTAuth
func (r *RiceHandler) DeleteRice(ctx *gin.Context) {
	id := ctx.Param("id")

	numID, err := strconv.Atoi(id)
	if err != nil {
		validationError(ctx, errors.New("id must be a number"))
		return
	}

	err = r.svc.DeleteRice(ctx, numID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, nil)
}
