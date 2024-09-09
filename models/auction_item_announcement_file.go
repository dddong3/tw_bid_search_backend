package models

import (
	"gorm.io/gorm"
)

type AuctionItemAnnouncementFile struct {
	gorm.Model
	ID           int       `gorm:"primaryKey"`
	Court        string    `gorm:"size:100;not null"`
	CaseYear     string    `gorm:"size:100;not null"`
	CaseID       string    `gorm:"size:100;not null"`
	CaseNo       string    `gorm:"size:100;not null"`
	FileType     string    `gorm:"size:100;not null"`
	AnnouncementFile []byte `gorm:"type:bytea;not null"`
}

func (a AuctionItemAnnouncementFile) TableName() string {
	return "AUCTION_ITEM_ANNOUNCEMENT"
}

