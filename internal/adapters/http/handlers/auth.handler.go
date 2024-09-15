package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type AuthHandler struct {
	svc ports.IAuthService
}

func NewAuthHandler(authService ports.IAuthService) *AuthHandler {
	return &AuthHandler{
		svc: authService,
	}
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"ex@mail.com" format:"email"`
	Password string `json:"password" binding:"required,min=8" example:"12345678" minLength:"8"`
}

// Login go-blog
//
//	@Summary		Login and get an access token
//	@Description	Logs in a registered user and returns an access token if the credentials are valid.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		loginRequest				true	"Login request body"
//	@Success		200		{object}	response{data=authResponse}	"Successfully logged in"
//	@Failure		400		{object}	errorResponse				"Validation error"
//	@Failure		401		{object}	errorResponse				"Unauthorized error"
//	@Failure		500		{object}	errorResponse				"Internal server error"
//	@Router			/auth/login [post]
func (auth AuthHandler) Login(ctx *gin.Context) {
	var req loginRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	token, err := auth.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		handleError(ctx, err)
		return
	}
	res := newAuthResponse(token)

	handleSuccess(ctx, res)
}
