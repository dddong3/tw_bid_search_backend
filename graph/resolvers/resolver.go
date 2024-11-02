package resolvers

import (
	"github.com/dddong3/Bid_Backend/auctionitem"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuctionItemService *auctionitem.AuctionItemService
}

func InitResolver() *Resolver {
	return &Resolver{
		AuctionItemService: &auctionitem.AuctionItemService{
			Repo: auctionitem.GetAuctionItemRepo(),
		},
	}
}
