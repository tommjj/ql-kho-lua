package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

// newResponse create a response body
func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

// meta represents metadata for a paginated response
type meta struct {
	Total int `json:"total" example:"100"`
	Limit int `json:"limit" example:"10"`
	Skip  int `json:"skip" example:"0"`
}

// newMeta is a helper function to create metadata for a paginated response
func newMeta(total, limit, skip int) meta {
	return meta{
		Total: total,
		Limit: limit,
		Skip:  skip,
	}
}

// authResponse type to auth response for auth handler
type authResponse struct {
	Token string `json:"token" example:"eyJJ9.eyJpEzNDR9.fUjDw0"`
}

// newAuthResponse create a auth response for login handler
func newAuthResponse(token string) authResponse {
	return authResponse{
		Token: token,
	}
}

// uploadImageResponse represents a upland image response body
type uploadImageResponse struct {
	Filename string `json:"filename" example:"name.ext"`
}

// newUploadImageResponse is a helper function to create a response body for handling upload image data
func newUploadImageResponse(filename string) uploadImageResponse {
	return uploadImageResponse{
		Filename: filename,
	}
}

// userResponse represents a user response body
type userResponse struct {
	ID    int         `json:"id" example:"1"`
	Name  string      `json:"name" example:"vertin"`
	Phone string      `json:"phone" example:"+84123456789"`
	Email string      `json:"email" example:"example@exm.com"`
	Role  domain.Role `json:"role" example:"member"`
}

// newUserResponse is a helper function to create a response body for handling user data
func newUserResponse(user *domain.User) userResponse {
	return userResponse{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
		Email: user.Email,
		Role:  user.Role,
	}
}

// storehouseResponse represents a storehouse response body
type storehouseResponse struct {
	ID       int       `json:"id"`
	Name     string    `json:"name" example:"store 01"`
	Location []float64 `json:"location" example:"[50.12,68.36]"`
	Image    string    `json:"image" example:"2455.png"`
	Capacity int       `json:"capacity" example:"1200"`
}

// newStorehouseResponse is a helper function to create a response body for handling storehouse data
func newStorehouseResponse(store *domain.Storehouse) storehouseResponse {
	latitude, longitude, _ := store.ParseLocation()

	return storehouseResponse{
		ID:       store.ID,
		Name:     store.Name,
		Location: []float64{latitude, longitude},
		Image:    store.Image,
		Capacity: store.Capacity,
	}
}

// listResponse represents a list with meta response items
type listResponse struct {
	Meta  meta `json:"meta"`
	Items any  `json:"items"`
}

// newListResponse is a helper function to create a response body for handling users data
func newListResponse(meta meta, items any) listResponse {
	return listResponse{
		Meta:  meta,
		Items: items,
	}
}

// errorStatusMap is a map of defined error messages and their corresponding http status codes
var errorStatusMap = map[error]int{
	domain.ErrInternal:                   http.StatusInternalServerError,
	domain.ErrDataNotFound:               http.StatusNotFound,
	domain.ErrConflictingData:            http.StatusConflict,
	domain.ErrInvalidCredentials:         http.StatusUnauthorized,
	domain.ErrUnauthorized:               http.StatusUnauthorized,
	domain.ErrEmptyAuthorizationHeader:   http.StatusUnauthorized,
	domain.ErrInvalidAuthorizationHeader: http.StatusUnauthorized,
	domain.ErrInvalidAuthorizationType:   http.StatusUnauthorized,
	domain.ErrInvalidToken:               http.StatusUnauthorized,
	domain.ErrExpiredToken:               http.StatusUnauthorized,
	domain.ErrForbidden:                  http.StatusForbidden,
	domain.ErrNoUpdatedData:              http.StatusBadRequest,
}

// handleSuccess write success response with status code 200 mess Success and data
func handleSuccess(ctx *gin.Context, data any) {
	res := newResponse(true, "Success", data)
	ctx.JSON(http.StatusOK, res)
}

// handleError write error response
func handleError(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := praseError(err)
	res := newErrorResponse(errMsg)
	ctx.JSON(statusCode, res)
}

// errorResponse type of error response
type errorResponse struct {
	Success  bool     `json:"success" example:"false"`
	Messages []string `json:"messages" example:"data not found"`
}

// newErrorResponse create an new error response
func newErrorResponse(errMegs []string) errorResponse {
	return errorResponse{
		Success:  false,
		Messages: errMegs,
	}
}

// validationError handle validation error, write err response
func validationError(ctx *gin.Context, err error) {
	errMegs := praseError(err)
	res := newErrorResponse(errMegs)

	ctx.JSON(http.StatusBadRequest, res)
}

// praseError prase error to error messages
func praseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, validationError := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, validationError.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}
	return errMsgs
}
