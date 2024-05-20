package processors

import (
	"context"
	"errors"

	"ethereumBlockchainParser/internal/parsers"

	"github.com/ethereum/go-ethereum/ethclient"
)

type ObserveAddressTransactionsProcessor struct {
	ctx            context.Context
	ethereumParser parsers.BasicParser
}

type ObserveAddressTransactionsRequest struct {
	Address string `json:"address"`
}

type ObserveAddressTransactionsResponse struct {
	DebugMessage   string `json:"debugMessage"`
	ObserveSuccess bool   `json:"observeSuccess"`
}

func NewObserveAddressTransactionsProcessor(ctx context.Context, client *ethclient.Client) *ObserveAddressTransactionsProcessor {
	return &ObserveAddressTransactionsProcessor{
		ctx:            ctx,
		ethereumParser: parsers.NewEthereumParser(ctx, client),
	}
}

func (p *ObserveAddressTransactionsProcessor) ObserveAddressTransactions(req *ObserveAddressTransactionsRequest) *ObserveAddressTransactionsResponse {
	err := validateObserveAddressTransactionsRequest(req)
	if err != nil {
		return &ObserveAddressTransactionsResponse{
			DebugMessage: err.Error(),
		}
	}

	observeSuccess, err := p.ethereumParser.Subscribe(req.Address)
	if err != nil {
		return &ObserveAddressTransactionsResponse{
			DebugMessage: err.Error(),
		}
	}

	return &ObserveAddressTransactionsResponse{
		ObserveSuccess: observeSuccess,
	}
}

func validateObserveAddressTransactionsRequest(request *ObserveAddressTransactionsRequest) error {
	if request == nil {
		return errors.New("request cannot be nil")
	}

	if len(request.Address) == 0 {
		return errors.New("request address cannot be empty")
	}

	return nil
}
