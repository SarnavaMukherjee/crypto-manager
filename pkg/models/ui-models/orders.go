/**
 * File: orders.go
 * Author: Sarnava Mukherjee
 * Contact: (sarnavamukherjee20@gmail.com)
 */

package db_models

import (
	"github.com/SarnavaMukherjee/crypto-manager/pkg/models/enums"
)

type Order struct {
	Order     enums.OrderType    `json:"order"`
	Coin      enums.CoinType     `json:"coin"`
	Holding   float64            `json:"holding"`
	Price     float64            `json:"price"`
	UpdatedOn string             `json:"updatedOn"`
}
