package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrthkc/peagym/internal/pkg/conf"
	"github.com/mrthkc/peagym/internal/pkg/service"
	log "github.com/sirupsen/logrus"
)

// Listen : starts api by listening to incoming requests
func Listen(c *conf.Config) {
	service.Config = c

	if service.Config.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(CORS())

	// router.go
	route(r)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server can not start: %v", err)
	}
}

// CORS : basic cors settings
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
