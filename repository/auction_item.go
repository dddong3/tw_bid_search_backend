package repository

import (
	"github.com/dddong3/Bid_Backend/models"
	"github.com/dddong3/Bid_Backend/logger"
)


func (r *AuctionItemRepo) GetFileTypes(court, caseYear, caseID, caseNo string) ([]string, error) {
	var fileTypes []string
	err := r.DB.Model(&models.AuctionItemAnnouncementFile{}).Where("court = ? AND case_year = ? AND case_id = ? AND case_no = ?", court, caseYear, caseID, caseNo).
		Pluck("file_type", &fileTypes).Error

	return fileTypes, err
}

func (r *AuctionItemRepo) GetPDF(court, caseYear, caseID, caseNo, fileType string) ([]byte, error) {
	var auctionItem models.AuctionItemAnnouncementFile
	err := r.DB.Where("court = ? AND case_year = ? AND case_id = ? AND case_no = ? AND file_type = ?", court, caseYear, caseID, caseNo, fileType).
		First(&auctionItem).Error
	return auctionItem.AnnouncementFile, err
}

func (r *AuctionItemRepo) GetAuctionItemsWithPage(limit int, page int) ([]*models.AuctionItem, int64, error) {
	var auctionItems []*models.AuctionItem
	var total int64
	err := r.DB.Model(&models.AuctionItem{}).Count(&total).Error
	if err != nil {
		logger.Logger.Errorf("Error counting auction items: %v", err)
		return nil, 0, err
	}

	err = r.DB.Order("id ASC").Offset(limit * (page - 1)).Limit(limit).Find(&auctionItems).Error
	return auctionItems, total, err
}

func (r *AuctionItemRepo) GetAuctionItemByID(id int) (*models.AuctionItem, error) {
	var auctionItem *models.AuctionItem
	err := r.DB.First(&auctionItem, id).Error
	return auctionItem, err
}