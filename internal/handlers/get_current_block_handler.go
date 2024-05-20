package handlers

import (
	"encoding/json"
	"net/http"

	"ethereumBlockchainParser/internal/processors"
)

type GetCurrentBlockHandler struct {
	Processor *processors.GetCurrentBlockProcessor
}

func NewGetCurrentBlockHandler(processor *processors.GetCurrentBlockProcessor) *GetCurrentBlockHandler {
	return &GetCurrentBlockHandler{processor}
}

func (h *GetCurrentBlockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := h.Processor.GetCurrentBlock()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
