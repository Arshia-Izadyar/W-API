package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"wapi/src/api/helper"
	"wapi/src/common"
	"wapi/src/config"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
	"wapi/src/services"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	service *services.FileService
}

func NewFileHandler(cfg *config.Config) *FileHandler {
	return &FileHandler{
		service: services.NewFileService(cfg),
	}
}

// CreateFile godoc
// @Summary Create File
// @Description Create File
// @Tags File
// @Accept x-www-form-urlencoded
// @produces json
// @Param file formData dto.UploadFileRequest true "CreateFile"
// @Param file formData file true "CreateFile"
// @Success 200 {object} helper.Response "CreateFile response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/files/create [post]
// @Security AuthBearer
func (fh *FileHandler) CreateFile(ctx *gin.Context) {
	upload := dto.UploadFileRequest{}
	err := ctx.ShouldBind(&upload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.ValidationError), err))
		return
	}
	req := dto.CreateFileRequest{}
	req.Description = upload.Description
	req.MineType = upload.File.Header.Get("Content-Type")
	req.Directory = "../../uploads"
	req.Name, err = common.SaveFile(upload.File, req.Directory)
	fmt.Println(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}
	res, err := fh.service.GenericCreateFile(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, int(helper.Success)))
}

// UpdateFile godoc
// @Summary Update File
// @Description Update File
// @Tags File
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateFileRequest true "UpdateFile"
// @Success 200 {object} helper.Response "UpdateFile response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/files/update/{id} [put]
// @Security AuthBearer
func (fh *FileHandler) UpdateFile(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	req := dto.UpdateFileRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, int(helper.ValidationError), err))
		return
	}
	res, err := fh.service.GenericUpdateFile(ctx, id, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, int(helper.Success)))

}

// GetFile godoc
// @Summary Get File
// @Description Get File
// @Tags File
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "GetFile response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/files/get/{id} [get]
// @Security AuthBearer
func (fh *FileHandler) GetFile(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponseWithError(nil, false, int(helper.NotFoundError), errors.New("not found id = 0")))
		return
	}

	res, err := fh.service.GenericGetFileById(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, int(helper.Success)))

}

// DeleteFile godoc
// @Summary Delete File
// @Description Delete File
// @Tags File
// @Accept json
// @produces json
// @Success 200 {object} helper.Response "DeleteFile response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/files/delete/{id} [delete]
// @Security AuthBearer
func (fh *FileHandler) DeleteFile(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponseWithError(nil, false, int(helper.NotFoundError), errors.New("not found id = 0")))
		return
	}

	file, _ := fh.service.GenericGetFileById(ctx, id)
	err := os.Remove(fmt.Sprintf("%s/%s", file.Directory, file.Name))

	if err != nil {
		logger.Error(err, logging.IO, logging.Delete, "cant delete file", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.ValidationError), err))
		return
	}

	err = fh.service.GenericDeleteFile(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}
	ctx.JSON(http.StatusNoContent, helper.GenerateBaseResponse(gin.H{"Status": "Deleted"}, true, int(helper.Success)))

}
