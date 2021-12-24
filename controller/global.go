package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
