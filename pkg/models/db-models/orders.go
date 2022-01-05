/**
 * File: orders.go
 * Author: Sarnava Mukherjee
 * Contact: (sarnavamukherjee20@gmail.com)
 */

package db_models

import (
	"github.com/SarnavaMukherjee/crypto-manager/pkg/models/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OderID    primitive.ObjectID `json:"orderID,omitempty" bson:"_id,omitempty"`
	Order     enums.OrderType    `json:"order"`
	Coin      enums.CoinType     `json:"coin"`
	Holding   float64            `json:"holding"`
	Price     float64            `json:"price"`
	Amount    float64            `json:"amount"`
	UpdatedOn uint64             `json:"updatedOn"`
}
