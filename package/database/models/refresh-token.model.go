package models

import (
	"time"
)

type RefreshToken struct {
	ID string `gorm:"primarykey; type:uuid; default:uuid_generate_v4()" json:"id"`

	UserID string `gorm:"not null" json:"user_id"`
	User   User   `gorm:"foreignKey:UserID" json:"-"`

	Token     string    `gorm:"not null;unique" json:"refresh_token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}
