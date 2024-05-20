package processors

import (
	"context"
	"errors"
	"ethereumBlockchainParser/internal/data"
	"ethereumBlockchainParser/internal/parsers"

	"github.com/ethereum/go-ethereum/ethclient"
)

type GetTransactionsProcessor struct {
	ctx            context.Context
	ethereumParser parsers.BasicParser
}

type GetTransactionsRequest struct {
	Address   string `json:"address"`
	FromBlock *int64 `json:"from_block"`
	ToBlock   *int64 `json:"to_block"`
}

type GetTransactionsResponse struct {
	DebugMessage string              `json:"debugMessage"`
	Transactions []*data.Transaction `json:"transactions"`
}

func NewGetTransactionsProcessor(ctx context.Context, client *ethclient.Client) *GetTransactionsProcessor {
	return &GetTransactionsProcessor{
		ctx:            ctx,
		ethereumParser: parsers.NewEthereumParser(ctx, client),
	}
}

func (p *GetTransactionsProcessor) GetTransactions(getTransactionsRequest *GetTransactionsRequest) *GetTransactionsResponse {
	err := validateGetTransactionsRequest(getTransactionsRequest)
	if err != nil {
		return &GetTransactionsResponse{
			DebugMessage: err.Error(),
		}
	}

	transactions, err := p.ethereumParser.GetTransactions(getTransactionsRequest.Address, getTransactionsRequest.FromBlock, getTransactionsRequest.ToBlock)
	if err != nil {
		return &GetTransactionsResponse{
			DebugMessage: err.Error(),
		}
	}

	return &GetTransactionsResponse{
		Transactions: transactions,
	}
}

func validateGetTransactionsRequest(request *GetTransactionsRequest) error {
	if request == nil {
		return errors.New("request is nil")
	}

	if len(request.Address) == 0 {
		return errors.New("request address is empty")
	}

	return nil
}
