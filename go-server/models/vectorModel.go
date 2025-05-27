package models

import "gorm.io/gorm"

type VectorModel struct {
	gorm.Model
	ID     string    `gorm:"primaryKey;uniqueIndex;not null" json:"id"`
	Vector []float64 `gorm:"type:float[]" json:"vector"`
}
