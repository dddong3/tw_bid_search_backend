package models

import (
	"gorm.io/gorm"
	"time"
)

type AuctionItem struct {
	gorm.Model
	ID           int       `gorm:"primaryKey"`
	FilePath     string    `gorm:"size:100;not null"`
	RowID        int       `gorm:"not null"`
	CaseYear     string    `gorm:"size:100;not null"`
	CaseID       string    `gorm:"size:100;not null"`
	CaseNo       string    `gorm:"size:100;not null"`
	SaleDate     time.Time `gorm:"not null"`
	SaleNo       int       `gorm:"not null"`
	Name         string    `gorm:"size:100;not null"`
	Quantity     string    `gorm:"size:100;default:''"`
	Unit         string    `gorm:"size:100;default:''"`
	Notes        string    `gorm:"type:text;default:''"`
	Remark       string    `gorm:"type:text;default:''"`
	Court        string    `gorm:"size:100;not null"`
	PicturePath  string    `gorm:"size:100;not null"`
	PictureCount int       `gorm:"default:0"`
	TotalPrice   int       `gorm:"default:0"`
	Deposit      string    `gorm:"size:100;default:''"`
}

func (a AuctionItem) TableName() string {
	return "AUCTION_ITEM"
}

