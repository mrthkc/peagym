package main

import (
	"flag"

	"github.com/mrthkc/peagym/internal/api"
	"github.com/mrthkc/peagym/internal/pkg/conf"
	log "github.com/sirupsen/logrus"
)

func main() {
	env := flag.String("env", "local", "environment")
	flag.Parse()

	config, err := conf.NewConfig("config/default.yml", *env)
	if err != nil {
		log.Fatalf("Can not reead config: %v", err)
	}

	config.Env = *env
	config.DBCred = config.Mysql.Local
	config.Secret = config.JWT.Local.Secret
	if config.Env == "prod" {
		config.DBCred = config.Mysql.Prod
		config.Secret = config.JWT.Prod.Secret
	}

	log.Info("SDK API started w/Env: " + *env)
	api.Listen(config)
}
