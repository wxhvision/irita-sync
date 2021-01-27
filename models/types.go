// interface for a document

package models

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	CollectionNameTxn = "sync_txn"
)

var (
	SyncTaskModel SyncTask
	BlockModel    Block
	TxModel       Tx

	Collections = []Docs{
		SyncTaskModel,
		BlockModel,
		TxModel,
	}
)

type (
	Docs interface {
		// collection name
		Name() string
		// ensure indexes
		EnsureIndexes()
		// primary key pair(used to find a unique record)
		PkKvPair() map[string]interface{}
	}

	Coin struct {
		Denom  string `bson:"denom"`
		Amount string `bson:"amount"`
	}

	Coins []*Coin
)


func ensureDocsIndexes() {
	if len(Collections) > 0 {
		for _, v := range Collections {
			v.EnsureIndexes()
		}
	}
}

func BuildDocCoins(coins sdk.Coins) []Coin {
	var (
		res []Coin
	)
	if len(coins) > 0 {
		for _, v := range coins {
			c := Coin{
				Denom:  v.Denom,
				Amount: v.Amount.String(),
			}
			res = append(res, c)
		}
	}

	return res
}

