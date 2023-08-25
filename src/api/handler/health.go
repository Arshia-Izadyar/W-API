package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok get")

}
