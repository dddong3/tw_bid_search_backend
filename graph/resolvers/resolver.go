package resolvers

import (
	// "gorm.io/gorm"
	"github.com/dddong3/Bid_Backend/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	// DB *gorm.DB
	AuctionItemService *services.AuctionItemService
}
