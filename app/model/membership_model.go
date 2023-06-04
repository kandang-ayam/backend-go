package model

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type Membership struct {
	ID         int        `json:"id"`
	MemberCode string     `json:"member_code"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Phone      int        `json:"phone"`
	BirthDay   string     `json:"birth_day"`
	Level      string     `json:"level"`
	Point      int        `json:"points"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

func (m *Membership) AfterFind(tx *gorm.DB) (err error) {
	m.BirthDay = strings.Split(m.BirthDay, "T")[0]

	return
}
