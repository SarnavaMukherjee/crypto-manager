/**
 * File: crypto-manager.go
 * Author: Kedar Kekan
 * Contact: (support@airavana.ai)
 * Copyright (c) 2020 - 2021 Airavana Inc.
 */

package crypto_manager

import (
	"net/http"

	"github.com/SarnavaMukherjee/crypto-manager/pkg/api/v1/crypto-manager/wallet"
	"github.com/gin-gonic/gin"
)

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Status": "UP",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	pms := r.Group("/cms")
	{
		wallet.ApplyRoutes(pms)
	}
}
