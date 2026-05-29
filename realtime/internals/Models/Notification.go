package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	ID             string         `gorm:"primaryKey;type:char(36)" json:"id"`
	Type           string         `gorm:"type:string;not null" json:"type"`
	NotifiableType string         `gorm:"type:string;not null" json:"notifiable_type"`
	NotifiableID   uint64         `gorm:"type:bigint unsigned;not null" json:"notifiable_id"`
	Data           string         `gorm:"type:text;not null" json:"data"`
	ReadAt         *time.Time     `json:"read_at"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

func (n *Notification) BeforeCreate(tx *gorm.DB) (err error) {
	if n.ID == "" {
		n.ID = uuid.New().String()
	}
	return
}
