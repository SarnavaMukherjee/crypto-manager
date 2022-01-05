/**
 * File: main.go
 * Author: Sarnava Mukherjee
 * Contact: (sarnavamukherjee20@gmail.com)
 */

package main

import (
	"github.com/SarnavaMukherjee/crypto-manager/pkg/api"
	"github.com/SarnavaMukherjee/crypto-manager/pkg/config"
	"github.com/SarnavaMukherjee/crypto-manager/pkg/db/mongodb"
	"github.com/SarnavaMukherjee/crypto-manager/pkg/logger"
	"github.com/SarnavaMukherjee/crypto-manager/pkg/server"
)

func main() {

	config.CreateConfig()
	cfg := config.GetConfig()

	logger.CreateLogger(cfg.LOG_LEVEL, "CMS")

	mongodb.NewMongoDB(cfg.MONGO_USER, cfg.MONGO_PASSWORD, cfg.MONGO_HOST)

	app := server.Setup(cfg.SERVER_PORT, cfg.SERVER_MODE)
	app.Use(config.Inject(&cfg))
	api.ApplyRoutes(app)
	server.Start(app)
}
