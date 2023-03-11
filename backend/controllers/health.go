package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct{}

func (h HealthController) Status(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Working!")
}
