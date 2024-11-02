package rest

import (
	"github.com/dddong3/Bid_Backend/rest/handlers"
	"github.com/go-chi/chi"
)

func RegisterRoutes(r *chi.Mux, auctionItemHandler *handlers.AuctionItemHandler) {
	r.Get("/api/files/pdf/{court}/{case_year}/{case_id}/{case_no}/types", auctionItemHandler.GetFileTypes)
	r.Get("/api/files/pdf/{court}/{case_year}/{case_id}/{case_no}/{file_type}", auctionItemHandler.GetPDF)
}
