package handlers

import (
	"errors"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

// pagination is a metadata for pagination
type pagination struct {
	TotalRecords int  `json:"total_records" example:"5"`
	LimitRecords int  `json:"limit_records" example:"5"`
	CurrentPage  int  `json:"current_page" example:"2"`
	TotalPages   int  `json:"total_pages" example:"10"`
	NextPage     *int `json:"next_page" example:"2"`
	PrevPage     *int `json:"prev_page" example:"1"`
}

// newPagination create a new pagination metadata for pagination response
func newPagination(totalAllRecords int64, totalRecords, limitRecords, currentPage int) *pagination {
	var nextPage *int
	var privPage *int

	totalPages := int(math.Ceil(float64(totalAllRecords) / float64(limitRecords)))

	if currentPage > 1 {
		privPage = newPtr(currentPage - 1)
	}

	if currentPage < totalPages {
		nextPage = newPtr(currentPage + 1)
	}

	return &pagination{
		TotalRecords: totalRecords,
		LimitRecords: limitRecords,
		CurrentPage:  currentPage,
		TotalPages:   totalPages,
		NextPage:     nextPage,
		PrevPage:     privPage,
	}
}

// response is a response body
type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

// responseWithPagination is a response body with pagination
type responseWithPagination struct {
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Success"`
	Pagination *pagination `json:"pagination"`
	Data       any         `json:"data"`
}

// newResponseWithPagination create a response body with pagination
func newResponseWithPagination(success bool, message string, pagination *pagination, data any) responseWithPagination {
	return responseWithPagination{
		Success:    success,
		Message:    message,
		Pagination: pagination,
		Data:       data,
	}
}

// newResponse create a response body
func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

// authResponse represents a auth response body
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

// warehouseResponse represents a warehouse response body
type warehouseResponse struct {
	ID       int       `json:"id" example:"1"`
	Name     string    `json:"name" example:"store 01"`
	Location []float64 `json:"location" example:"50.12,68.36"`
	Image    string    `json:"image" example:"2455.png"`
	Capacity int       `json:"capacity" example:"1200"`
}

// newWarehouseResponse is a helper function to create a response body for handling warehouse data
func newWarehouseResponse(store *domain.Warehouse) warehouseResponse {
	latitude, longitude, _ := store.ParseLocation()

	return warehouseResponse{
		ID:       store.ID,
		Name:     store.Name,
		Location: []float64{latitude, longitude},
		Image:    store.Image,
		Capacity: store.Capacity,
	}
}

// usedCapacityResponse represents a used capacity response data
type usedCapacityResponse struct {
	UsedCapacity int64 `json:"used_capacity" example:"500"`
}

// newUsedCapacityResponse is a helper function to create a response body for handling used capacity data
func newUsedCapacityResponse(v int64) usedCapacityResponse {
	return usedCapacityResponse{
		UsedCapacity: v,
	}
}

// warehouseItemResponse represents a item in warehouse
type warehouseItemResponse struct {
	ID       int    `json:"id" example:"1"`
	RiceName string `json:"rice_name" example:"name"`
	Capacity int    `json:"capacity" example:"500"`
}

// newWarehouseItemResponse is a helper function to create a response body for handling warehouse item data
func newWarehouseItemResponse(v *domain.WarehouseItem) warehouseItemResponse {
	w := warehouseItemResponse{
		ID:       v.RiceID,
		Capacity: v.Quantity,
	}

	if v.Rice != nil {
		w.RiceName = v.Rice.Name
	}

	return w
}

// riceResponse represents a rice response body
type riceResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// newRiceResponse is a helper function to create a response body for handling rice data
func newRiceResponse(rice *domain.Rice) riceResponse {
	return riceResponse{
		ID:   rice.ID,
		Name: rice.Name,
	}
}

// customerResponse represents a customer response body
type customerResponse struct {
	ID      int    `json:"id" example:"1"`
	Name    string `json:"name" example:"Ascalon"`
	Email   string `json:"email" example:"ascalon@exp.com"`
	Phone   string `json:"phone" example:"+84123456789"`
	Address string `json:"address" example:"abc, eyz"`
}

// newCustomerResponse is a helper function to create a response body for handling customer data
func newCustomerResponse(customer *domain.Customer) customerResponse {
	return customerResponse{
		ID:      customer.ID,
		Name:    customer.Name,
		Email:   customer.Email,
		Phone:   customer.Phone,
		Address: customer.Address,
	}
}

// invoiceDetailResponse represents a invoice detail response body
type invoiceDetailResponse struct {
	RiceID   int     `json:"rice_id" example:"1"`
	Name     string  `json:"name" example:"name"`
	Price    float64 `json:"price" example:"500"`
	Quantity int     `json:"quantity" example:"5"`
}

// NewInvoiceDetail is a helper function to create a invoice Detail response for handling invoice data
func newInvoiceDetail(invoiceDetail *domain.InvoiceItem) invoiceDetailResponse {
	res := invoiceDetailResponse{
		RiceID:   invoiceDetail.RiceID,
		Price:    invoiceDetail.Price,
		Quantity: invoiceDetail.Quantity,
	}

	if invoiceDetail.Rice != nil {
		res.Name = invoiceDetail.Rice.Name
	}
	return res
}

// invoiceResponse is a helper function to create a response body for handling invoice data
type invoiceResponse struct {
	ID            int                     `json:"id" example:"1"`
	CustomerID    int                     `json:"customer_id" example:"1"`
	CustomerName  string                  `json:"customer_name,omitempty" example:"Ascalon"`
	WarehouseID   int                     `json:"warehouse_id" example:"1"`
	WarehouseName string                  `json:"warehouse_name,omitempty" example:"store 01"`
	UserID        int                     `json:"user_id" example:"1"`
	UserName      string                  `json:"user_name,omitempty" example:"vertin"`
	CreatedAt     time.Time               `json:"created_at" example:"2021-09-01T00:00:00Z"`
	TotalPrice    float64                 `json:"total_price" example:"500"`
	Details       []invoiceDetailResponse `json:"details,omitempty"`
}

// newInvoiceResponse is a helper function to create a invoice response for handling invoice data
func newInvoiceResponse(invoice *domain.Invoice) invoiceResponse {
	res := invoiceResponse{
		ID:          invoice.ID,
		CustomerID:  invoice.CustomerID,
		WarehouseID: invoice.WarehouseID,
		UserID:      invoice.UserID,
		CreatedAt:   invoice.CreatedAt,
		TotalPrice:  invoice.TotalPrice,
		Details:     make([]invoiceDetailResponse, 0, len(invoice.Details)),
	}

	if invoice.CreatedBy != nil {
		res.UserName = invoice.CreatedBy.Name
	}
	if invoice.Customer != nil {
		res.CustomerName = invoice.Customer.Name
	}
	if invoice.Warehouse != nil {
		res.WarehouseName = invoice.Warehouse.Name
	}

	for _, v := range invoice.Details {
		res.Details = append(res.Details, newInvoiceDetail(&v))
	}
	return res
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
	domain.ErrWarehouseFull:              http.StatusBadRequest,
	domain.ErrInsufficientStock:          http.StatusBadRequest,
}

// handleSuccess write success response with status code 200 mess Success and data
func handleSuccessPagination(ctx *gin.Context, pagination *pagination, data any) {
	res := newResponseWithPagination(true, "Success", pagination, data)
	ctx.JSON(http.StatusOK, res)
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
