package model

import "time"

type Product struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Image      string    `json:"image"`
	Code       string    `json:"code"`
	CategoryID int       `json:"category_id"`
	Quantity   int       `json:"quantity"`
	Unit       string    `json:"unit"`
	Price      int       `json:"price"`
	Category   Category  `json:"category" gorm:"foreignKey:CategoryID"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
