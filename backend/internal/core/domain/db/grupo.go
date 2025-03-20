package db

import "time"

type Leal struct {
	ID            uint      `gorm:"primaryKey"`
	Nombre        uint      `gorm:"not null;"`
	FechaCreacion time.Time `gorm:"autoCreateTime"`
}

func (Leal) TableName() string {
	return "leal"
}
