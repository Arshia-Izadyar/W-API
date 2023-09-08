package handler

import (
	"net/http"
	"wapi/src/api/helper"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck godoc
// @Summary Health Check
// @Description Health Check
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.Response "Success"
// @Failure 400 {object} helper.Response "Failed"
// @Router /v1/health/ [get]
func (h *HealthHandler) HealthGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("ok", true, http.StatusOK))
}
