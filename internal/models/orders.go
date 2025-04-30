package models

import "time"

type Orders struct {
	ID              uint
	UserId          uint
	PayableAmount   float64
	CurrentStatus   string
	PackedStatus    bool
	ShippedStatus   bool
	ShippedDate     time.Time
	DeliveredStatus bool
	DeliveredDate   time.Time
	CancelledStatus bool
	CancelledDate   time.Time
	ReturnedStatus  bool
	ReturnedDate    time.Time
	PaymentMode     string
	Reason          string
	Items           []OrderItems
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type OrderItems struct {
	ProductId    uint
	Quantity     int
	Images       []string
	SkuCode      string
	Name         string
	Desceription string
	SellingPrice float64
	Volume       int
	Length       int
	Width        int
	Height       int
	Brand        string
	Category     string
}
