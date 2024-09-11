package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/dddong3/Bid_Backend/graph"
	"github.com/dddong3/Bid_Backend/logger"
	"github.com/dddong3/Bid_Backend/models"
)

// GetAuctionItems is the resolver for the getAuctionItems field.
func (r *queryResolver) GetAuctionItems(ctx context.Context, page *int, limit *int) (*graph.AuctionItemConnection, error) {
	const defaultPage, defaultLimit = 1, 10
	if page == nil {
		page = new(int)
		*page = defaultPage
	}
	if limit == nil {
		limit = new(int)
		*limit = defaultLimit
	}
	if *page < defaultPage {
		*page = defaultPage
	}
	if *limit < 1 {
		*limit = defaultLimit
	}

	auctionItems, hasNextPage, hasPrevPage, totalCount, err := r.AuctionItemService.GetAuctionItemsWithPage(*limit, *page)

	if err != nil {
		logger.Logger.Errorf("Error fetching auction items: %v", err)
		return nil, err
	}

	logger.Logger.Debugf("Fetched %d auction items", len(auctionItems))

	return &graph.AuctionItemConnection{
		Nodes: auctionItems,
		PageInfo: &graph.PageInfo{
			HasNextPage: hasNextPage,
			HasPrevPage: hasPrevPage,
			TotalCount:  totalCount,
		},
	}, nil
}

// GetAuctionItemWithID is the resolver for the getAuctionItemWithId field.
func (r *queryResolver) GetAuctionItemWithID(ctx context.Context, id *int) (*graph.SingleAuctionItem, error) {
	if id == nil {
		logger.Logger.Error("id is required")
		return nil, fmt.Errorf("id is required")
	}

	logger.Logger.Debugf("Fetched auction item with id %d", *id)

	item, err := r.AuctionItemService.GetAuctionItemByID(*id)

	if err != nil {
		logger.Logger.Errorf("Error fetching auction item: %v", err)
		return nil, err
	}

	return &graph.SingleAuctionItem{
		Node: item,
	}, nil
}

// GetAuctionItemWithRelate is the resolver for the getAuctionItemWithRelate field.
func (r *queryResolver) GetAuctionItemWithRelate(ctx context.Context, court *string, year *string, caseID *string, caseNo *string) ([]*models.AuctionItem, error) {
	if court == nil || year == nil || caseID == nil || caseNo == nil {
		logger.Logger.Error("court, year, caseID, caseNo are required")
		return nil, fmt.Errorf("court, year, caseID, caseNo are required")
	}

	logger.Logger.Debugf("Fetched auction item with court %s, year %s, caseID %s, caseNo %s", *court, *year, *caseID, *caseNo)

	items, err := r.AuctionItemService.GetAuctionItemWithRelate(*court, *year, *caseID, *caseNo)

	if err != nil {
		logger.Logger.Errorf("Error fetching auction item: %v", err)
		return nil, err
	}

	return items, nil
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
