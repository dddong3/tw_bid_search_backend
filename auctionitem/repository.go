package auctionitem

import (
	"fmt"
	"strings"
	"time"

	"github.com/dddong3/Bid_Backend/database"
	"github.com/dddong3/Bid_Backend/logger"

	"gorm.io/gorm"
)

type AuctionItemRepo struct {
	DB *gorm.DB
}

var auctionItemRepoInstance *AuctionItemRepo

func GetAuctionItemRepo() *AuctionItemRepo {
	if auctionItemRepoInstance == nil {
		db := database.GetDB()
		db.AutoMigrate(&AuctionItem{})
		db.AutoMigrate(&AuctionItemAnnouncementFile{})
		auctionItemRepoInstance = &AuctionItemRepo{DB: db}
	}
	return auctionItemRepoInstance
}

func (r *AuctionItemRepo) GetFileTypes(court, caseYear, caseID, caseNo string) ([]string, error) {
	var fileTypes []string
	err := r.DB.Model(&AuctionItemAnnouncementFile{}).Where("court = ? AND case_year = ? AND case_id = ? AND case_no = ?", court, caseYear, caseID, caseNo).
		Pluck("file_type", &fileTypes).Error

	return fileTypes, err
}

func (r *AuctionItemRepo) GetPDF(court, caseYear, caseID, caseNo, fileType string) ([]byte, error) {
	var auctionItem AuctionItemAnnouncementFile
	err := r.DB.Where("court = ? AND case_year = ? AND case_id = ? AND case_no = ? AND file_type = ?", court, caseYear, caseID, caseNo, fileType).
		First(&auctionItem).Error
	return auctionItem.AnnouncementFile, err
}

func (r *AuctionItemRepo) GetAuctionItemsWithPage(limit int, page int) ([]*AuctionItem, int64, error) {
	var auctionItems []*AuctionItem
	var total int64
	err := r.DB.Model(&AuctionItem{}).Count(&total).Error
	if err != nil {
		logger.Logger.Errorf("Error counting auction items: %v", err)
		return nil, 0, err
	}

	err = r.DB.Order("id ASC").Offset(limit * (page - 1)).Limit(limit).Find(&auctionItems).Error
	return auctionItems, total, err
}

func (r *AuctionItemRepo) GetAuctionItemByID(id int) (*AuctionItem, error) {
	var auctionItem *AuctionItem
	err := r.DB.First(&auctionItem, id).Error
	return auctionItem, err
}

func (r *AuctionItemRepo) GetAuctionItemWithRelate(court, year, caseID, caseNo string) ([]*AuctionItem, error) {
	var auctionItems []*AuctionItem
	err := r.DB.Where("court = ? AND case_year = ? AND case_id = ? AND case_no = ?", court, year, caseID, caseNo).Find(&auctionItems).Error
	return auctionItems, err
}

func (r *AuctionItemRepo) GetAuctionItemsWithQuery(limit, page int, targetEmbedding []float32, startDate, endDate time.Time) ([]*AuctionItem, int64, error) {
	var (
		auctionItems []*AuctionItem
		total        int64
	)
	// embeddingStr := "ARRAY["
	// for i, val := range targetEmbedding {
	// 	if i > 0 {
	// 		embeddingStr += ", "
	// 	}
	// 	embeddingStr += fmt.Sprintf("%f", val)
	// }
	// embeddingStr += "]::vector"
	embeddingStr := "ARRAY[" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(targetEmbedding)), ","), "[]") + "]::vector"

	// query := fmt.Sprintf(`
	// 	SELECT * FROM "AUCTION_ITEM"
	// 	WHERE "START_DATE" >= '%s' AND "END_DATE" <= '%s'
	// 	ORDER BY EMBEDDING <=> %s
	// 	LIMIT %d OFFSET %d
	// `, embeddingStr, limit, limit*(page-1))

	query := r.DB.Table("AUCTION_ITEM").
		Where("sale_date >= ? AND sale_date <= ?", startDate, endDate).
		Order(fmt.Sprintf("EMBEDDING <=> %s", embeddingStr)).
		Limit(limit).
		Offset(limit * (page - 1))

	// err := r.DB.Raw(query).Scan(&auctionItems).Error
	err := query.Find(&auctionItems).Error
	if err != nil {
		logger.Logger.Errorf("Error fetching auction items: %v", err)
		return nil, 0, err
	}

	// err = r.DB.Model(&AuctionItem{}).Count(&total).Error
	err = query.Count(&total).Error
	if err != nil {
		logger.Logger.Errorf("Error counting auction items: %v", err)
		return nil, 0, err
	}

	return auctionItems, total, nil

}
