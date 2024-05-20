package data

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

type Transaction struct {
	Hash        string
	From        string
	To          string
	Value       *big.Int
	BlockNumber uint64
	Gas         uint64
	GasPrice    *big.Int
	Nonce       uint64
}

func NewTransaction(tx *types.Transaction, txReceipt *types.Receipt, msg *core.Message) *Transaction {
	if tx == nil || txReceipt == nil || msg == nil {
		return nil
	}

	return &Transaction{
		Hash:        tx.Hash().Hex(),
		From:        msg.From.String(),
		To:          msg.To.String(),
		Value:       tx.Value(),
		BlockNumber: txReceipt.BlockNumber.Uint64(),
		Gas:         tx.Gas(),
		GasPrice:    tx.GasPrice(),
		Nonce:       tx.Nonce(),
	}
}
