package handler

import (
	"net/http"
	"wapi/src/api/helper"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("ok", true, http.StatusOK))

}
