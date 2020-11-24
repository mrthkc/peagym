package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mrthkc/peagym/internal/pkg/service"
)

func route(r *gin.Engine) {
	// Base
	r.GET("/", service.BaseHandler)
}
