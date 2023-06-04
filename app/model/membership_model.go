package model

import "time"

type Membership struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     int       `json:"phone"`
	BirthDay  time.Time `json:"birth_day"`
	Level     string    `json:"level"`
	Point     int       `json:"point"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
