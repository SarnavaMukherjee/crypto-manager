/**
 * File: order-helper.go
 * Author: Sarnava Mukherjee
 * Contact: (support@airavana.ai)
 * Copyright (c) 2020 - 2021 Airavana Inc.
 */

package helper

import (
	"github.com/SarnavaMukherjee/crypto-manager/pkg/db/mongodb"
	log "github.com/SarnavaMukherjee/crypto-manager/pkg/logger"
	dbModels "github.com/SarnavaMukherjee/crypto-manager/pkg/models/db-models"
	"github.com/SarnavaMukherjee/crypto-manager/pkg/models/enums"
	uiModels "github.com/SarnavaMukherjee/crypto-manager/pkg/models/ui-models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func AddNewOrder(orderRequest uiModels.Order) (primitive.ObjectID, error) {

	updatedOn, err := time.Parse(time.RFC3339, orderRequest.UpdatedOn)
	if err != nil {
		log.Error("[ADD-ORDER-FAIL] failed to add order, error: %s", err.Error())
		return primitive.NilObjectID, err
	}

	order := dbModels.Order{
		Order: orderRequest.Order,
		Coin: orderRequest.Coin,
		Holding: orderRequest.Holding,
		Price: orderRequest.Price,
		UpdatedOn: uint64(updatedOn.Unix()),
	}

	if order.UpdatedOn == 0 {
		order.UpdatedOn = uint64(time.Now().Unix())
	}

	order.Amount = order.Holding * order.Price

	if order.Order == enums.BUY {
		return AddNewBuyOrder(order)
	} else if order.Order == enums.SELL {
		return AddNewSellOrder(order)
	}

	return primitive.NilObjectID, nil
}

func AddNewBuyOrder(order dbModels.Order) (primitive.ObjectID, error) {

	coinDetails, err := mongodb.GetCoinSummeryFromWallet(order.Coin)
	if err != nil {
		log.Error("[UPDATING-WALLET-FAIL] failed to get coin details from wallet, error: %s", err.Error())
		if err := mongodb.UpdateCoinSummeryFromWallet(coinDetails); err != nil {
			log.Error("[UPDATING-WALLET-FAIL] failed to update coin details yo wallet, error: %s", err.Error())
			return primitive.NilObjectID, err
		}
	} else {
		oldAvgPrice := coinDetails.AvgPrice

		coinDetails.AvgPrice = (oldAvgPrice*coinDetails.Holding + order.Holding*order.Price) / (coinDetails.Holding + order.Holding)
		coinDetails.Holding += order.Holding
		coinDetails.AmountInvested = coinDetails.AvgPrice * coinDetails.Holding

		if err := mongodb.UpdateCoinSummeryFromWallet(coinDetails); err != nil {
			log.Error("[UPDATING-WALLET-FAIL] failed to update coin details to wallet, error: %s", err.Error())
			return primitive.NilObjectID, err
		}
	}

	orderID, err := mongodb.AddNewOrder(order)
	if err != nil {
		log.Error("[ADD-ORDER-FAIL] failed to add order, error: %s", err.Error())
		return primitive.NilObjectID, err
	}

	log.Info("[ADD-ORDER] order added, orderID: %s", orderID)

	return orderID, nil
}

func AddNewSellOrder(order dbModels.Order) (primitive.ObjectID, error) {
	coinDetails, err := mongodb.GetCoinSummeryFromWallet(order.Coin)
	if err != nil {
		log.Error("[UPDATING-WALLET-FAIL] failed to get coin details from wallet, error: %s", err.Error())
		return primitive.NilObjectID, nil
	}

	coinDetails.Holding -= order.Holding

	if coinDetails.Holding == 0 {
		coinDetails.AvgPrice = 0
	}
	coinDetails.AmountInvested = 0

	if err := mongodb.UpdateCoinSummeryFromWallet(coinDetails); err != nil {
		log.Error("[UPDATING-WALLET-FAIL] failed to update coin details to wallet, error: %s", err.Error())
		return primitive.NilObjectID, err
	}

	orderID, err := mongodb.AddNewOrder(order)
	if err != nil {
		log.Error("[ADD-ORDER-FAIL] failed to add order, error: %s", err.Error())
		return primitive.NilObjectID, err
	}

	log.Info("[ADD-ORDER] order added, orderID: %s", orderID)

	return orderID, nil
}

func GetAllOrders(coinType, orderType string) ([]uiModels.Order, error) {

	orders, err := mongodb.GetAllOrders(orderType, coinType)
	if err != nil {
		log.Error("[ADD-ORDER-FAIL] failed to add order, error: %s", err.Error())
		return []uiModels.Order{}, err
	}

	var uiOrder []uiModels.Order
	for _, order := range orders {
		newUIOrder := uiModels.Order{
			Order: order.Order,
			Coin:      order.Coin,
			Holding: order.Holding,
			Price:order.Price,
			UpdatedOn: time.Unix(int64(order.UpdatedOn), 0).String(),
		}

		uiOrder = append(uiOrder, newUIOrder)
	}


	return uiOrder, nil
}

func GetWalletDetails(coinType string) ([]dbModels.Wallet, error) {

	wallet, err := mongodb.GetWalletDetails(coinType)
	if err != nil {
		log.Error("[GET-WALLET-DETAILS] failed to add order, error: %s", err.Error())
		return []dbModels.Wallet{}, err
	}

	return wallet, nil
}