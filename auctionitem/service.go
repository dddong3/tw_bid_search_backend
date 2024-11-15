package auctionitem

import (
	"context"
	"github.com/dddong3/Bid_Backend/config"
	"github.com/dddong3/Bid_Backend/logger"
	openai "github.com/sashabaranov/go-openai"
	"time"
)

type AuctionItemService struct {
	Repo *AuctionItemRepo
}

func (s *AuctionItemService) GetFileTypes(court, caseYear, caseID, caseNo string) ([]string, error) {
	return s.Repo.GetFileTypes(court, caseYear, caseID, caseNo)
}

func (s *AuctionItemService) GetPDF(court, caseYear, caseID, caseNo, fileType string) ([]byte, error) {
	return s.Repo.GetPDF(court, caseYear, caseID, caseNo, fileType)
}

func (s *AuctionItemService) GetAuctionItemsWithPage(limit *int, page *int) ([]*AuctionItem, bool, bool, int, error) {
	const defaultPage, defaultLimit = 1, 10
	if page == nil {
		page = new(int)
		*page = defaultPage
	}
	if *page < defaultPage {
		*page = defaultPage
	}
	if limit == nil {
		limit = new(int)
		*limit = defaultLimit
	}
	if *limit < 1 {
		*limit = defaultLimit
	}
	logger.Logger.Debugf("Fetching auction items with limit: %d, page: %d", *limit, *page)
	items, total, err := s.Repo.GetAuctionItemsWithPage(*limit, *page)
	if err != nil {
		logger.Logger.Errorf("Error fetching auction items: %v", err)
		return nil, false, false, 0, err
	}
	hasNextPage := (*page * *limit) < int(total)
	hasPrevPage := *page > 1
	return items, hasNextPage, hasPrevPage, int(total), nil
}

func (s *AuctionItemService) GetAuctionItemByID(id int) (*AuctionItem, error) {
	return s.Repo.GetAuctionItemByID(id)
}

func (s *AuctionItemService) GetAuctionItemsWithQuery(query string, startData time.Time, endDate time.Time, limit, page int) ([]*AuctionItem, bool, bool, int, error) {
	const defaultPage, defaultLimit = 1, 10
	if page < defaultPage {
		page = defaultPage
	}
	if limit < 1 {
		limit = defaultLimit
	}
	logger.Logger.Debugf("Fetching auction items with query: %s, limit: %d, page: %d", query, limit, page)

	var similarityThreshold float32 = 0.6

	if query == "" {
		similarityThreshold = 1
	}

	client := openai.NewClient(config.GetEnv("OPENAI_API_KEY", ""))

	queryParams := &openai.EmbeddingRequest{
		Model: openai.SmallEmbedding3,
		Input: &query,
	}

	resp, err := client.CreateEmbeddings(context.Background(), queryParams)

	if err != nil {
		logger.Logger.Errorf("Error fetching auction items: %v", err)
		return nil, false, false, 0, err
	}

	items, total, err := s.Repo.GetAuctionItemsWithQuery(limit, page, resp.Data[0].Embedding, startData, endDate, similarityThreshold)

	if err != nil {
		logger.Logger.Errorf("Error fetching auction items: %v", err)
		return nil, false, false, 0, err
	}
	hasNextPage := (page * limit) < int(total)
	hasPrevPage := page > 1
	return items, hasNextPage, hasPrevPage, int(total), nil
}

func (s *AuctionItemService) GetAuctionItemWithRelate(court, year, caseID, caseNo string) ([]*AuctionItem, error) {
	return s.Repo.GetAuctionItemWithRelate(court, year, caseID, caseNo)
}
