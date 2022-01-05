/**
 * File: db-operations.go
 * Author: Sarnava Mukherjee
 * Contact: (support@airavana.ai)
 * Copyright (c) 2020 - 2021 Airavana Inc.
 */

package mongodb

import (
	"context"
	dbModels "github.com/SarnavaMukherjee/crypto-manager/pkg/models/db-models"
	"github.com/SarnavaMukherjee/crypto-manager/pkg/models/enums"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ORDERS  = "orders"
	WALLETS = "wallet"
)

func AddNewOrder(orders dbModels.Order) (primitive.ObjectID, error) {

	var mongoService MongoService
	mongoService.BindCollection("crypto", ORDERS)

	result, err := mongoService.Collection.InsertOne(context.TODO(), orders)
	return result.InsertedID.(primitive.ObjectID), err
}

func GetCoinSummeryFromWallet(coinType enums.CoinType) (dbModels.Wallet, error) {
	var mongoService MongoService
	mongoService.BindCollection("crypto", WALLETS)

	var coinDetails dbModels.Wallet
	filterOptions := bson.D{{"coin", coinType}}

	err := mongoService.Collection.FindOne(context.TODO(), filterOptions).Decode(&coinDetails)
	return coinDetails, err
}

func UpdateCoinSummeryFromWallet(coin dbModels.Wallet) error {
	var mongoService MongoService
	mongoService.BindCollection("crypto", WALLETS)

	filterOptions := bson.D{{"coin", coin.Coin}}

	updateOptions := options.Update()
	updateOptions.SetUpsert(true)

	_, err := mongoService.Collection.UpdateOne(context.TODO(), filterOptions, bson.D{{"$set", coin}}, updateOptions)
	return err
}

func GetAllOrders(order, coinType string) ([]dbModels.Order, error) {
	var mongoService MongoService
	mongoService.BindCollection("crypto", ORDERS)
	filter := bson.D{}
	if len(order) > 0 {
		filter = append(filter, bson.E{"order", order})
	}

	if len(coinType) > 0 {
		filter = append(filter, bson.E{"coin", coinType})
	}

	cur, err := mongoService.Collection.Find(context.TODO(), filter)
	if err != nil {
		return []dbModels.Order{}, err
	}

	var orders []dbModels.Order
	err = cur.All(context.TODO(), &orders)
	if err != nil {
		return []dbModels.Order{}, err
	}

	return orders, nil
}

func GetWalletDetails(coinType string) ([]dbModels.Wallet, error) {
	var mongoService MongoService
	mongoService.BindCollection("crypto", WALLETS)

	filter := bson.D{}
	if len(coinType) > 0 {
		filter = append(filter, bson.E{"coin", coinType})
	}

	cur, err := mongoService.Collection.Find(context.TODO(), filter)
	if err != nil {
		return []dbModels.Wallet{}, err
	}

	var wallet []dbModels.Wallet
	err = cur.All(context.TODO(), &wallet)
	if err != nil {
		return []dbModels.Wallet{}, err
	}

	return wallet, nil
}
