package handlers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

var (
	// authorizationHeaderKey is the key for authorization header in the request
	authorizationHeaderKey = "authorization"
	// authorizationType is the accepted authorization type
	authorizationType = "jwt"
	// authorizationPayloadKey is the key for authorization payload in the context
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(token ports.ITokenService) gin.HandlerFunc {
	v := validator.New()
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		isEmpty := len(authorizationHeader) == 0

		if isEmpty {
			handleError(ctx, domain.ErrInvalidAuthorizationHeader)
			ctx.Abort()
			return
		}

		fields := strings.Fields(authorizationHeader)
		isValid := len(fields) == 2
		if !isValid {
			handleError(ctx, domain.ErrInvalidAuthorizationHeader)
			ctx.Abort()
			return
		}

		currentAuthorizationType := strings.ToLower(fields[0])
		if currentAuthorizationType != authorizationType {
			handleError(ctx, domain.ErrInvalidAuthorizationType)
			ctx.Abort()
			return
		}

		accessToken := fields[1]
		err := v.Var(accessToken, "required,jwt")
		if err != nil {
			handleError(ctx, domain.ErrInvalidToken)
			ctx.Abort()
			return
		}

		payload, err := token.VerifyToken(accessToken)
		if err != nil {
			handleError(ctx, err)
			ctx.Abort()
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

// RoleRootMiddleware is a middleware to check if the user is a root
func RoleRootMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := getAuthPayload(ctx, authorizationPayloadKey)

		isRoot := token.Role == domain.Root

		if !isRoot {
			handleError(ctx, domain.ErrForbidden)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
