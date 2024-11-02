package auctionitem

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"time"
)

type Vector []float32

func (v *Vector) Scan(src interface{}) error {
	str, ok := src.(string)
	if !ok {
		return errors.New("unsupported Scan type for Vector")
	}
	var vec []float32
	if err := json.Unmarshal([]byte(str), &vec); err != nil {
		return err
	}
	*v = vec
	return nil
}

func (v Vector) Value() (driver.Value, error) {
	return json.Marshal(v)
}

type AuctionItem struct {
	gorm.Model
	ID           int       `gorm:"primaryKey"`
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
	UpdatedAt    time.Time `gorm:"not null"`
	Embedding    Vector    `gorm:"type:vector(1536);"`
}

func (a AuctionItem) TableName() string {
	return "AUCTION_ITEM"
}

type AuctionItemAnnouncementFile struct {
	gorm.Model
	ID               int       `gorm:"primaryKey"`
	Court            string    `gorm:"size:100;not null"`
	CaseYear         string    `gorm:"size:100;not null"`
	CaseID           string    `gorm:"size:100;not null"`
	CaseNo           string    `gorm:"size:100;not null"`
	FileType         string    `gorm:"size:100;not null"`
	AnnouncementFile []byte    `gorm:"type:bytea;not null"`
	UpdatedAt        time.Time `gorm:"not null"`
}

func (a AuctionItemAnnouncementFile) TableName() string {
	return "AUCTION_ITEM_ANNOUNCEMENT"
}
