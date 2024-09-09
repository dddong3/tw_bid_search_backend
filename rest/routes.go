package rest

import (
	"github.com/go-chi/chi"
	"github.com/dddong3/Bid_Backend/rest/handlers"
)

func RegisterRoutes(r *chi.Mux, auctionItemHandler *handlers.AuctionItemHandler) {
	r.Get("/api/files/pdf/{court}/{case_year}/{case_id}/{case_no}/types", auctionItemHandler.GetFileTypes)
	r.Get("/api/files/pdf/{court}/{case_year}/{case_id}/{case_no}/{file_type}", auctionItemHandler.GetPDF)
}
