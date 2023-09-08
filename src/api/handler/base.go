package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"wapi/src/api/helper"
	"wapi/src/config"
	"wapi/src/dto"
	"wapi/src/pkg/logging"

	"github.com/gin-gonic/gin"
)

var logger = logging.NewLogger(config.LoadCfg())

// input 	output
func Create[Ti any, To any](ctx *gin.Context, input func(ctx context.Context, req *Ti) (res *To, err error)) {
	req := new(Ti)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, int(helper.ValidationError), err))
		return
	}
	res, err := input(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, int(helper.Success)))
}

func Update[Ti, To any](ctx *gin.Context, caller func(c context.Context, id int, req *Ti) (res *To, err error)) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	req := new(Ti)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, int(helper.ValidationError), err))
		return
	}
	res, err := caller(ctx, id, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, int(helper.Success)))
}

func Delete(ctx *gin.Context, caller func(c context.Context, id int) error) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponseWithError(nil, false, int(helper.NotFoundError), errors.New("not found id = 0")))
		return
	}
	err := caller(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}
	ctx.JSON(http.StatusNoContent, helper.GenerateBaseResponse(gin.H{"Status": "Deleted"}, true, int(helper.Success)))

}

func GetById[To any](ctx *gin.Context, caller func(ctx context.Context, id int) (*To, error)) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponseWithError(nil, false, int(helper.NotFoundError), errors.New("not found id = 0")))
		return
	}
	// req := new(To)

	// err := ctx.ShouldBindJSON(&req)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, int(helper.ValidationError), err))
	// 	return
	// }
	res, err := caller(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}

func GetByFilter[Ti, To any](ctx *gin.Context, caller func(ctx context.Context, req *Ti) (*dto.PageList[To], error)) {
	req := new(Ti)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.ValidationError), err))
		return
	}
	res, err := caller(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, int(helper.Success)))
}
