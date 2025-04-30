package models

import "time"

type Products struct {
	ID             uint      `gorm:"primarykey" json:"id"`
	Name           string    `gorm:"not null" json:"name"`
	SkuCode        string    `gorm:"not null;unique" json:"skuCode"`
	Description    string    `gorm:"not null" json:"description"`
	Images         []string  `gorm:"type:text[]" json:"images"`
	Quantity       int       `gorm:"not null" json:"quantity"`
	CostPrice      float64   `gorm:"not null" json:"costPrice"`
	SellingPrice   float64   `gorm:"not null" json:"sellingPrice"`
	MainPrice      float64   `gorm:"not null" json:"mainPrice"`
	Gst            float64   `gorm:"not null" json:"gst"`
	Status         string    `gorm:"not null" json:"status"`
	Volume         int       `gorm:"not null" json:"volume"`
	Length         int       `gorm:"not null" json:"length"`
	Width          int       `gorm:"not null" json:"width"`
	Height         int       `gorm:"not null" json:"height"`
	Brand          string    `gorm:"not null" json:"brand"`
	Category       string    `gorm:"not null" json:"category"`
	IsDiscountable string    `gorm:"not null" json:"isDiscountable"`
	ExpiryDate     string    `gorm:"not null" json:"expiryDate"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	UpdatedBy      string    `gorm:"-" json:"updatedBy,omitempty"`
	DeletedAt      time.Time `gorm:"index" json:"deletedAt,omitempty"`
	DeletedBy      string    `gorm:"-" json:"deletedBy,omitempty"`
}

type ProductDetails struct {
	ID           uint
	Name         string
	SellingPrice float64
	SkuCode      string
	Quantity     int
	Description  string
	Images       []string
	Brand        string
}
