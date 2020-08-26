package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"name": "index",
	})
}
