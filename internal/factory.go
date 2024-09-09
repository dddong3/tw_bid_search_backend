package internal

import (
	"gorm.io/gorm"
	"github.com/dddong3/Bid_Backend/repository"
	"github.com/dddong3/Bid_Backend/services"
	"github.com/dddong3/Bid_Backend/graph/resolvers"
)


func InitAuctionItemService(db *gorm.DB) *services.AuctionItemService {
	auctionRepo := repository.NewAuctionItemRepo(db)
	auctionService := services.NewAuctionItemService(auctionRepo)
	return auctionService
}


func InitResolver(db *gorm.DB) *resolvers.Resolver {
	auctionService := InitAuctionItemService(db)
	return &resolvers.Resolver{
		AuctionItemService: auctionService,
	}
}