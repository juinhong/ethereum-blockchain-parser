package parsers

import (
	"context"
	"ethereumBlockchainParser/internal/data"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumParser struct {
	BasicParser

	ctx    context.Context
	client *ethclient.Client
}

func NewEthereumParser(ctx context.Context, client *ethclient.Client) BasicParser {
	return &EthereumParser{
		ctx:    ctx,
		client: client,
	}
}

func (p *EthereumParser) GetCurrentBlock() (int64, error) {
	header, err := p.client.HeaderByNumber(p.ctx, nil)
	if err != nil {
		return 0, err
	}

	blockNumber := header.Number
	block, err := p.client.BlockByNumber(p.ctx, blockNumber)
	if err != nil {
		return 0, err
	}

	return block.Number().Int64(), nil
}

func (p *EthereumParser) Subscribe(address string) (bool, error) {
	addr := common.HexToAddress(address)

	q := ethereum.FilterQuery{
		Addresses: []common.Address{addr},
	}
	logCh := make(chan types.Log)
	sub, err := p.client.SubscribeFilterLogs(p.ctx, q, logCh)
	if err != nil {
		return false, err
	}
	defer sub.Unsubscribe()

	go func() {
		for {
			select {
			case err := <-sub.Err():
				fmt.Printf("[ERROR] SubscribeFilterLogs subscription failed|err: %v", err)
				return
			case lg := <-logCh:
				fmt.Println("New log received:", lg)
			}
		}
	}()

	return true, nil
}

func (p *EthereumParser) GetTransactions(address string, fromBlock, toBlock *int64) ([]*data.Transaction, error) {
	addr := common.HexToAddress(address)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{addr},
	}
	if fromBlock != nil {
		query.FromBlock = big.NewInt(*fromBlock)
	}
	if toBlock != nil {
		query.ToBlock = big.NewInt(*toBlock)
	}

	logs, err := p.client.FilterLogs(p.ctx, query)
	if err != nil {
		return nil, err
	}

	var transactions []*data.Transaction
	for _, lg := range logs {
		tx, isPending, err := p.client.TransactionByHash(p.ctx, lg.TxHash)
		if err != nil {
			log.Printf("[ERROR] TransactionByHash library call failed|err: %v", err)
			continue
		}

		if isPending {
			continue
		}

		receipt, err := p.client.TransactionReceipt(p.ctx, lg.TxHash)

		msg, err := core.TransactionToMessage(tx, types.LatestSignerForChainID(tx.ChainId()), tx.GasPrice())
		if err != nil {
			log.Printf("[ERROR] TransactionToMessage library call failed|err: %v", err)
			continue
		}

		transaction := data.NewTransaction(tx, receipt, msg)
		if transaction != nil {
			transactions = append(transactions, transaction)
		}
	}

	return transactions, nil
}
