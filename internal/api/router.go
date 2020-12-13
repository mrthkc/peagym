package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mrthkc/peagym/internal/pkg/service"
)

func route(r *gin.Engine) {
	// Base
	r.GET("/api/", service.BaseHandler)

	// Register
	r.POST("/api/user", service.Register)
	r.POST("/api/login", service.Login)

	authorized := r.Group("/api/")
	authorized.Use(service.JWTAuth())
	{
		authorized.GET("token", service.Token)
		authorized.GET("user/:uid/profile", service.Profile)
	}
}
