package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type UserHandler struct {
	svc ports.IUserRepository
}

func NewUserHandler(userService ports.IUserService) *UserHandler {
	return &UserHandler{
		svc: userService,
	}
}

type createUserRequest struct {
	Name     string `json:"name" binding:"required,min=3" example:"vertin"`
	Email    string `json:"email" binding:"required,email" example:"example@exm.com"`
	Phone    string `json:"phone" binding:"required,e164" example:"+84123456788"`
	Password string `json:"password" binding:"required,min=8" example:"password"`
}

// CreateUser ql-kho-lua
//
//	@Summary		Create a new user and get created user data
//	@Description	Create a new user and get created user data
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		createUserRequest			true	"Create user body"
//	@Success		200		{object}	response{data=userResponse}	"Created user data"
//	@Failure		400		{object}	errorResponse				"Validation error"
//	@Failure		401		{object}	errorResponse				"Unauthorized error"
//	@Failure		403		{object}	errorResponse				"Forbidden error"
//	@Failure		409		{object}	errorResponse				"Conflicting data error"
//	@Failure		500		{object}	errorResponse				"Internal server error"
//	@Router			/users  [post]
//	@Security		JWTAuth
func (u *UserHandler) CreateUser(ctx *gin.Context) {
	var req createUserRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	createdUser, err := u.svc.CreateUser(ctx, &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
		Role:     domain.Member,
	})
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newUserResponse(createdUser)
	handleSuccess(ctx, res)
}

// GetUserByID ql-kho-lua
//
//	@Summary		Get a user
//	@Description	Get a user by user id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int							true	"User id"
//	@Success		200	{object}	response{data=userResponse}	"User data"
//	@Failure		400	{object}	errorResponse				"Validation error"
//	@Failure		401	{object}	errorResponse				"Unauthorized error"
//	@Failure		403	{object}	errorResponse				"Forbidden error"
//	@Failure		404	{object}	errorResponse				"Data not found error"
//	@Failure		500	{object}	errorResponse				"Internal server error"
//	@Router			/users/{id}  [get]
//	@Security		JWTAuth
func (u *UserHandler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	numID, err := strconv.Atoi(id)
	if err != nil {
		validationError(ctx, errors.New("id must be a number"))
		return
	}

	token := getAuthPayload(ctx, authorizationPayloadKey)

	isRootUser := token.Role == domain.Root
	if !isRootUser {
		if token.ID != numID {
			handleError(ctx, domain.ErrForbidden)
			return
		}
	}

	user, err := u.svc.GetUserByID(ctx, numID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newUserResponse(user)
	handleSuccess(ctx, res)
}

type getListUserRequest struct {
	Query string `form:"q" binding:"" example:"teo"`
	Skip  int    `form:"skip" binding:"min=0" example:"0"`
	Limit int    `form:"limit" binding:"min=5" example:"5"`
}

// GetListUsers ql-kho-lua
//
//	@Summary		get users
//	@Description	get users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			q		query		string										false	"Query"
//	@Param			skip	query		int											false	"Skip"	default(1)	minimum(1)
//	@Param			limit	query		int											false	"Limit"	default(5)	minimum(5)
//	@Success		200		{object}	responseWithPagination{data[]userResponse}	"Users data"
//	@Failure		400		{object}	errorResponse								"Validation error"
//	@Failure		401		{object}	errorResponse								"Unauthorized error"
//	@Failure		403		{object}	errorResponse								"Forbidden error"
//	@Failure		404		{object}	errorResponse								"Data not found error"
//	@Failure		500		{object}	errorResponse								"Internal server error"
//	@Router			/users [get]
//	@Security		JWTAuth
func (u *UserHandler) GetListUsers(ctx *gin.Context) {
	req := getListUserRequest{
		Skip:  1,
		Limit: 5,
	}
	err := ctx.BindQuery(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	count, err := u.svc.CountUsers(ctx, req.Query)
	if err != nil {
		handleError(ctx, err)
		return
	}

	if checkPageOverflow(count, req.Limit, req.Skip) {
		handleError(ctx, domain.ErrDataNotFound)
		return
	}

	users, err := u.svc.GetListUsers(ctx, req.Query, req.Limit, req.Skip)
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := make([]userResponse, 0, len(users))
	for _, user := range users {
		res = append(res, newUserResponse(&user))
	}

	pagination := newPagination(count, len(users), req.Limit, req.Skip)
	handleSuccessPagination(ctx, pagination, res)
}

type updateUserRequest struct {
	Name     string `json:"name" binding:"omitempty,min=3" example:"vertin"`
	Email    string `json:"email" binding:"omitempty,email" example:"example@exm.com"`
	Phone    string `json:"phone" binding:"omitempty,e164" example:"+84123456788"`
	Password string `json:"password" binding:"omitempty,min=8" example:"password"`
}

// UpdateUser ql-kho-lua
//
//	@Summary		update user
//	@Description	update user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"User id"
//	@Param			request	body		updateUserRequest			true	"Update user body"
//	@Success		200		{object}	response{data=userResponse}	"Updated user data"
//	@Failure		400		{object}	errorResponse				"Validation error"
//	@Failure		401		{object}	errorResponse				"Unauthorized error"
//	@Failure		403		{object}	errorResponse				"Forbidden error"
//	@Failure		404		{object}	errorResponse				"Data not found error"
//	@Failure		409		{object}	errorResponse				"Conflicting data error"
//	@Failure		500		{object}	errorResponse				"Internal server error"
//	@Router			/users/{id} [patch]
//	@Security		JWTAuth
func (u *UserHandler) UpdateUser(ctx *gin.Context) {
	var req updateUserRequest

	id := ctx.Param("id")

	numID, err := strconv.Atoi(id)
	if err != nil {
		validationError(ctx, errors.New("id must be a number"))
		return
	}

	token := getAuthPayload(ctx, authorizationPayloadKey)

	isRootUser := token.Role == domain.Root
	if !isRootUser {
		if token.ID != numID {
			handleError(ctx, domain.ErrForbidden)
			return
		}
	}

	err = ctx.BindJSON(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	isReqEmpty := req.Email == "" && req.Name == "" && req.Password == "" && req.Phone == ""
	if isReqEmpty {
		handleError(ctx, domain.ErrNoUpdatedData)
		return
	}

	updatedUser, err := u.svc.UpdateUser(ctx, &domain.User{
		ID:       numID,
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newUserResponse(updatedUser)
	handleSuccess(ctx, res)
}

// DeleteUserByID ql-kho-lua
//
//	@Summary		delete user
//	@Description	delete user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int				true	"User id"
//	@Success		200	{object}	response		"Deleted data"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		401	{object}	errorResponse	"Unauthorized error"
//	@Failure		403	{object}	errorResponse	"Forbidden error"
//	@Failure		404	{object}	errorResponse	"Data not found error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/users/{id} [delete]
//	@Security		JWTAuth
func (u *UserHandler) DeleteUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	numID, err := strconv.Atoi(id)
	if err != nil {
		validationError(ctx, errors.New("id must be a number"))
		return
	}

	user, err := u.svc.GetUserByID(ctx, numID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	if user.Role == domain.Root {
		handleError(ctx, domain.ErrForbidden)
		return
	}

	err = u.svc.DeleteUser(ctx, numID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, nil)
}
