package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

// getAuthPayload is a helper function to get the auth payload from the context
func getAuthPayload(ctx *gin.Context, key string) *domain.TokenPayload {
	return ctx.MustGet(key).(*domain.TokenPayload)
}

// newPtr return new pointer
func newPtr[T any](v T) *T {
	return &v
}

func checkPageOverflow(count int64, limit, skip int) bool {
	start := limit * (skip - 1)

	return count < int64(start)
}
