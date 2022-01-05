/**
 * File:  custom-scan-controller.go
 * Author: Kedar Kekan
 * Contact: (support@airavana.ai)
 * Copyright (c) 2020 - 2021 Airavana Inc.
 */

package wallet

import (
	"github.com/SarnavaMukherjee/crypto-manager/pkg/helper"
	log "github.com/SarnavaMukherjee/crypto-manager/pkg/logger"
	uiModels "github.com/SarnavaMukherjee/crypto-manager/pkg/models/ui-models"
	"github.com/SarnavaMukherjee/crypto-manager/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddNewOrder(ctx *gin.Context) {

	var order uiModels.Order
	if err := ctx.BindJSON(&order); err != nil {
		log.Error("", "", "[ADD-ORDER] failed to parse input payload, error: %s", err.Error())
		server.WriteResponseWithError(ctx, server.NewAppError(http.StatusBadRequest, "invalid input JSON payload"))
		return
	}

	orderID, err := helper.AddNewOrder(order)
	if err != nil {
		log.Error("[ADD-ORDER] failed to add new order, error: %s", err.Error())
		server.WriteResponseWithError(ctx, server.NewAppError(http.StatusInternalServerError, "failed to add new order"))
		return
	}

	response := map[string]interface{}{"orderID": orderID}

	server.WriteResponse(ctx, server.NewAppResponse(http.StatusOK, response, "new order added"))

}

func GetAllOrders(ctx *gin.Context) {

	coinType, ok := ctx.GetQuery("coin")
	if !ok {
		log.Error("[GET-ALL-ORDER] coin type missing")
	}

	orderType, ok := ctx.GetQuery("order")
	if !ok {
		log.Error("[GET-ALL-ORDER] order type missing")
	}

	orders, err := helper.GetAllOrders(coinType, orderType)
	if err != nil {
		log.Error("[GET-ALL-ORDER] failed to get all orders, error: %s", err.Error())
		server.WriteResponseWithError(ctx, server.NewAppError(http.StatusInternalServerError, "failed to get all orders"))
		return
	}

	response := map[string]interface{}{"dataList": orders, "total": len(orders)}
	server.WriteResponse(ctx, server.NewAppResponse(http.StatusOK, response, "found all scan histories"))
}

func GetWalletDetails(ctx *gin.Context) {

	coinType, ok := ctx.GetQuery("coin")
	if !ok {
		log.Error("[GET-ALL-ORDER] coin type missing")
	}

	orders, err := helper.GetWalletDetails(coinType)
	if err != nil {
		log.Error("[GET-ALL-ORDER] failed to get all orders, error: %s", err.Error())
		server.WriteResponseWithError(ctx, server.NewAppError(http.StatusInternalServerError, "failed to get all orders"))
		return
	}

	response := map[string]interface{}{"dataList": orders, "total": len(orders)}
	server.WriteResponse(ctx, server.NewAppResponse(http.StatusOK, response, "found all scan histories"))
}
