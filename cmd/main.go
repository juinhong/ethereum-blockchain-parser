package main

import (
	"context"
	"ethereumBlockchainParser/internal/handlers"
	"ethereumBlockchainParser/internal/processors"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()
	ethClient, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil || ethClient == nil {
		log.Fatalf("failed to connect to eth client: %v", err)
	}

	getCurrentBlockProcessor := processors.NewGetCurrentBlockProcessor(ctx, ethClient)
	getCurrentBlockHandler := handlers.NewGetCurrentBlockHandler(getCurrentBlockProcessor)
	http.Handle("/get_current_block", getCurrentBlockHandler)

	getTransactionsProcessor := processors.NewGetTransactionsProcessor(ctx, ethClient)
	getTransactionsHandler := handlers.NewGetTransactionsHandler(getTransactionsProcessor)
	http.Handle("/get_address_transactions", getTransactionsHandler)

	observeAddressTransactionsProcessor := processors.NewObserveAddressTransactionsProcessor(ctx, ethClient)
	observeAddressTransactionsHandler := handlers.NewObserveAddressTransactionsHandler(observeAddressTransactionsProcessor)
	http.Handle("/observe_address_transactions", observeAddressTransactionsHandler)

	_ = http.ListenAndServe(":8080", nil)
}
