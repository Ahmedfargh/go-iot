package models

import "gorm.io/gorm"

type Sector struct {
	gorm.Model
	Name        string   `gorm:"column:name;type:varchar(100);not null;unique" json:"name"`
	Description string   `gorm:"column:description;type:text" json:"description"`
	ParentID    *uint    `gorm:"column:parent_id" json:"parent_id"`
	Parent      *Sector  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children    []Sector `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	Devices     []Device `gorm:"foreignKey:SectorID" json:"devices,omitempty"`
}
