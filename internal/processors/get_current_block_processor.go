package processors

import (
	"context"
	"ethereumBlockchainParser/internal/parsers"

	"github.com/ethereum/go-ethereum/ethclient"
)

type GetCurrentBlockProcessor struct {
	ctx            context.Context
	ethereumParser parsers.BasicParser
}

type GetCurrentBlockResponse struct {
	DebugMessage string `json:"debugMessage"`
	BlockNumber  int64  `json:"block_number"`
}

func NewGetCurrentBlockProcessor(ctx context.Context, client *ethclient.Client) *GetCurrentBlockProcessor {
	return &GetCurrentBlockProcessor{
		ctx:            ctx,
		ethereumParser: parsers.NewEthereumParser(ctx, client),
	}
}

func (p *GetCurrentBlockProcessor) GetCurrentBlock() *GetCurrentBlockResponse {
	blockNumber, err := p.ethereumParser.GetCurrentBlock()
	if err != nil {
		return &GetCurrentBlockResponse{
			DebugMessage: err.Error(),
		}
	}

	return &GetCurrentBlockResponse{
		BlockNumber: blockNumber,
	}
}
