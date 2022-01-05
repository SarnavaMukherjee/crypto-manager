/**
 * File: wallet.go
 * Author: Sarnava Mukherjee
 * Contact: (support@airavana.ai)
 * Copyright (c) 2020 - 2021 Airavana Inc.
 */

package db_models

import "github.com/SarnavaMukherjee/crypto-manager/pkg/models/enums"

type Wallet struct {
	Coin           enums.CoinType `json:"coin"`
	Holding        float64        `json:"holding"`
	AvgPrice       float64        `json:"avgPrice"`
	AmountInvested float64        `json:"amountInvested"`
}
