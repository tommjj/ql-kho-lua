package handlers

import (
	"mime/multipart"
	"path/filepath"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

var imageExt = []string{".png", ".jpg", ".jpeg", ".jfif", ".gif"}

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

// UploadImage ql-kho-lua
//
//	@Summary		Upload image and get a file name
//	@Description	Upload temporary images for creating other resources
//	@Tags			upload
//	@Accept			mpfd
//	@Produce		json
//	@Param			request	body		uploadImageRequest					true	"Image file"
//	@Success		200		{object}	response{data=uploadImageResponse}	"Uploaded"
//	@Failure		400		{object}	errorResponse						"Validation error"
//	@Failure		401		{object}	errorResponse						"Unauthorized error"
//	@Failure		500		{object}	errorResponse						"Internal server error"
//	@Router			/upload [post]
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
