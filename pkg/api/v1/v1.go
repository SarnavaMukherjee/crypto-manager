/**
 * File: v1.go
 * Author: Kedar Kekan
 * Contact: (support@airavana.ai)
 * Copyright (c) 2020 - 2021 Airavana Inc.
 */

package api

import (
	cryptoManager "github.com/SarnavaMukherjee/crypto-manager/pkg/api/v1/crypto-manager"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		cryptoManager.ApplyRoutes(v1)
	}
}
