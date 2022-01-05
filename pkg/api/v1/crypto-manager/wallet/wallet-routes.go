/**
 * File: custom-scan-routes.go
 * Author: Kedar Kekan
 * Contact: (support@airavana.ai)
 * Copyright (c) 2020 - 2021 Airavana Inc.
 */

package wallet

import "github.com/gin-gonic/gin"

// ApplyRoutes ...
func ApplyRoutes(r *gin.RouterGroup) {

	wallet := r.Group("/wallet")
	{
		wallet.POST("/orders", AddNewOrder)

		wallet.GET("/orders", GetAllOrders)

		wallet.GET("", GetWalletDetails)
	}
}
