package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mrthkc/peagym/internal/pkg/conf"
	log "github.com/sirupsen/logrus"
)

// Config : config to all service
var Config *conf.Config

// BaseHandler : home - health-test
func BaseHandler(c *gin.Context) {
	log.Info("Base")
}
