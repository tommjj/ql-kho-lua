package handlers

import (
	"mime/multipart"
	"path/filepath"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

var imageExt = []string{".png", ".jpg", ".jpeg"}

type UploadHandler struct {
	svc ports.IUploadService
}

func NewUploadHandler(uploadService ports.IUploadService) *UploadHandler {
	return &UploadHandler{
		svc: uploadService,
	}
}

type uploadImageRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func (u UploadHandler) UploadImage(ctx *gin.Context) {
	var req uploadImageRequest

	err := ctx.Bind(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	filename := req.File.Filename
	ext := filepath.Ext(filename)

	isImage := slices.Contains(imageExt, ext)
	if !isImage {
		validationError(ctx, domain.ErrInvalidFileExt)
		return
	}

	newFilename, err := u.svc.SaveTemp(req.File)
	if err != nil {
		handleError(ctx, domain.ErrInternal)
	}

	res := newUploadImageResponse(newFilename)
	handleSuccess(ctx, res)
}
