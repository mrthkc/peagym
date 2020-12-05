package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mrthkc/peagym/internal/pkg/service"
)

func route(r *gin.Engine) {
	// Base
	r.GET("/", service.BaseHandler)

	// Register
	r.POST("/user", service.Register)
	r.POST("/login", service.Login)

	authorized := r.Group("/")
	authorized.Use(service.JWTAuth())
	{
		authorized.GET("/token", service.Token)
	}
}
