package model

import "time"

type Order struct {
	ID          int          `json:"id"`
	OrderID     string       `json:"order_id"`
	Name        string       `json:"name"`
	OrderOption string       `json:"order_option"`
	NumberTable int          `json:"number_table"`
	Note        string       `json:"note"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Items       []OrderItems `json:"items" gorm:"foreignKey:OrderID"`
	Transaction Transaction  `json:"transaction"`
}
