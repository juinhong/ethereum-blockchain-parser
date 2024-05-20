package handlers

import (
	"encoding/json"
	"net/http"

	"ethereumBlockchainParser/internal/processors"
)

type GetTransactionsHandler struct {
	Processor *processors.GetTransactionsProcessor
}

func NewGetTransactionsHandler(processor *processors.GetTransactionsProcessor) *GetTransactionsHandler {
	return &GetTransactionsHandler{processor}
}

func (handler *GetTransactionsHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	req := &processors.GetTransactionsRequest{}
	err := json.NewDecoder(request.Body).Decode(req)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	resp := handler.Processor.GetTransactions(req)
	response.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(response).Encode(resp)
}
