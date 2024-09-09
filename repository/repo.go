package repository

import (
	"gorm.io/gorm"
)

type AuctionItemRepo struct {
	DB *gorm.DB
}


func NewAuctionItemRepo(db *gorm.DB) *AuctionItemRepo {
	return &AuctionItemRepo{DB: db}
}