package models

import "gorm.io/gorm"

type Device struct {
	gorm.Model
	Identifier  string  `gorm:"column:identifier;type:varchar(50);not null;unique" json:"identifier"`
	Status      string  `gorm:"column:status;type:varchar(20);not null;default:'offline'" json:"status"`
	Temperature float64 `gorm:"column:temperature;type:decimal(5,2)" json:"temperature"`
	Humidity    float64 `gorm:"column:humidity;type:decimal(5,2)" json:"humidity"`
	SectorID    uint    `gorm:"column:sector_id" json:"sector_id"`
	Sector      *Sector `gorm:"foreignKey:SectorID" json:"sector,omitempty"`
}
