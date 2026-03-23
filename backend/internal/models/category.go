package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string     `gorm:"type:varchar(100);not null;uniqueIndex" json:"name"`
	Description string     `gorm:"type:text"                              json:"description"`
	Color       string     `gorm:"type:varchar(7);default:'#000000'"       json:"color"` // hex color for UI
	Questions   []Question `gorm:"foreignKey:CategoryID"                  json:"questions,omitempty"`
}
