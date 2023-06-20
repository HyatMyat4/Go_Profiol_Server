package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get_AllSkill() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	}
}
