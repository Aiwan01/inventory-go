package models

import "time"

type Users struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	FirstName string `gorm:"not null" validate:"required, min=2, max=50, firstName"`
	LastName  string `gorm:"not null" validate:"required, min=2, max=50, lastName"`
	Mobile    string `gorm:"not null" validate:"required, min=10, max=10, mobile"`
	Otp       int64
	Status    string
	Role      string
	CreatedAt time.Time `gorm:"json:createdAt"`
	UpdatedAt time.Time `gorm:"json:updatedAt"`
}

// just for basic have one table user, otherwise for customer different, and seller different and admin/office people have user collection
