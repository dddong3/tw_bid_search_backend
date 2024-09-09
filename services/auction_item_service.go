package services

import (
	"github.com/dddong3/Bid_Backend/repository"
	"github.com/dddong3/Bid_Backend/models"
)

type AuctionItemService struct {
	Repo *repository.AuctionItemRepo
}

func NewAuctionItemService(repo *repository.AuctionItemRepo) *AuctionItemService {
	return &AuctionItemService{Repo: repo}
}

func (s *AuctionItemService) GetFileTypes(court, caseYear, caseID, caseNo string) ([]string, error) {
	return s.Repo.GetFileTypes(court, caseYear, caseID, caseNo)
}

func (s *AuctionItemService) GetPDF(court, caseYear, caseID, caseNo, fileType string) ([]byte, error) {
	return s.Repo.GetPDF(court, caseYear, caseID, caseNo, fileType)
}

func (s *AuctionItemService) GetAuctionItemsWithPage(limit int, page int) ([]*models.AuctionItem, bool, bool, int, error) {
	items, total, err := s.Repo.GetAuctionItemsWithPage(limit, page)
	if err != nil {
		return nil, false, false, 0, err
	}
	hasNextPage := (page * limit) < int(total)
	hasPrevPage := page > 1
	return items, hasNextPage, hasPrevPage, int(total), nil
}

func (s *AuctionItemService) GetAuctionItemByID(id int) (*models.AuctionItem, error) {
	return s.Repo.GetAuctionItemByID(id)
}