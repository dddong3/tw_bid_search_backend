package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dddong3/Bid_Backend/logger"
	"github.com/dddong3/Bid_Backend/services"
	"github.com/go-chi/chi"
)

type AuctionItemHandler struct {
	Service *services.AuctionItemService
}

func (h *AuctionItemHandler) GetFileTypes(w http.ResponseWriter, r *http.Request) {
	// court := r.URL.Query().Get("court")
	// caseYear := r.URL.Query().Get("case_year")
	// caseID := r.URL.Query().Get("case_id")
	// caseNo := r.URL.Query().Get("case_no")
	court := chi.URLParam(r, "court")
	caseYear := chi.URLParam(r, "case_year")
	caseID := chi.URLParam(r, "case_id")
	caseNo := chi.URLParam(r, "case_no")


	logger.Logger.Debugf("court: %s, caseYear: %s, caseID: %s, caseNo: %s", court, caseYear, caseID, caseNo)


	fileTypes, err := h.Service.GetFileTypes(court, caseYear, caseID, caseNo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fileTypes)
}

func (h *AuctionItemHandler) GetPDF(w http.ResponseWriter, r *http.Request) {
	court := chi.URLParam(r, "court")
	caseYear := chi.URLParam(r, "case_year")
	caseID := chi.URLParam(r, "case_id")
	caseNo := chi.URLParam(r, "case_no")
	fileType := chi.URLParam(r, "file_type")

	pdf, err := h.Service.GetPDF(court, caseYear, caseID, caseNo, fileType)
	if err != nil {
		http.Error(w, "PDF not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Write(pdf)
}