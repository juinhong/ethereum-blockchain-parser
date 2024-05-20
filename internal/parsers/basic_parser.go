package parsers

import "ethereumBlockchainParser/internal/data"

type BasicParser interface {
	GetCurrentBlock() (int64, error)
	Subscribe(address string) (bool, error)
	GetTransactions(address string, fromBlock, toBlock *int64) ([]*data.Transaction, error)
}
