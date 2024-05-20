package handlers

import (
	"encoding/json"
	"net/http"

	"ethereumBlockchainParser/internal/processors"
)

type ObserveAddressTransactionsHandler struct {
	Processor *processors.ObserveAddressTransactionsProcessor
}

func NewObserveAddressTransactionsHandler(processor *processors.ObserveAddressTransactionsProcessor) *ObserveAddressTransactionsHandler {
	return &ObserveAddressTransactionsHandler{
		Processor: processor,
	}
}

func (handler *ObserveAddressTransactionsHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	req := &processors.ObserveAddressTransactionsRequest{}
	err := json.NewDecoder(request.Body).Decode(req)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	resp := handler.Processor.ObserveAddressTransactions(req)
	response.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(response).Encode(resp)
}
