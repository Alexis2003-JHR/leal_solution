package db

import (
	"time"
)

// Usuario
type User struct {
	ID           uint          `gorm:"primaryKey"`
	Name         string        `gorm:"not null"`
	Email        string        `gorm:"unique;not null"`
	CreatedAt    time.Time     `gorm:"autoCreateTime"`
	Transactions []Transaction `gorm:"foreignKey:UserID"`
	Redemptions  []Redemption  `gorm:"foreignKey:UserID"`
}

// Comercio
type Business struct {
	ID                uint               `gorm:"primaryKey"`
	Name              string             `gorm:"unique;not null"`
	TaxID             int                `gorm:"unique;not null"`
	Phone             int                `gorm:"not null"`
	Email             string             `gorm:"unique;not null"`
	CreatedAt         time.Time          `gorm:"autoCreateTime"`
	Branches          []Branch           `gorm:"foreignKey:BusinessID"`
	ConversionFactors []ConversionFactor `gorm:"foreignKey:BusinessID"`
	Rewards           []Reward           `gorm:"foreignKey:BusinessID"`
}

// Sucursal
type Branch struct {
	ID           uint          `gorm:"primaryKey"`
	BusinessID   uint          `gorm:"not null"`
	Name         string        `gorm:"not null"`
	CreatedAt    time.Time     `gorm:"autoCreateTime"`
	Business     Business      `gorm:"foreignKey:BusinessID"`
	Campaigns    []Campaign    `gorm:"foreignKey:BranchID"`
	Transactions []Transaction `gorm:"foreignKey:BranchID"`
}

// Factor de Conversión
type ConversionFactor struct {
	ID                  uint     `gorm:"primaryKey"`
	BusinessID          uint     `gorm:"not null"`
	Business            Business `gorm:"foreignKey:BusinessID"`
	PointsPerCurrency   float64  `gorm:"not null"`
	CashbackPerCurrency float64  `gorm:"not null"`
}

// Compra
type Transaction struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
	BranchID  uint      `gorm:"not null"`
	Branch    Branch    `gorm:"foreignKey:BranchID"`
	Amount    float64   `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Earnings  *Earnings `gorm:"foreignKey:TransactionID"`
}

// Campaña
type Campaign struct {
	ID                 uint      `gorm:"primaryKey"`
	BranchID           uint      `gorm:"not null"`
	Branch             Branch    `gorm:"foreignKey:BranchID"`
	StartDate          time.Time `gorm:"not null"`
	EndDate            time.Time `gorm:"not null"`
	PointsMultiplier   float64   `gorm:"default:1.0"`
	CashbackMultiplier float64   `gorm:"default:1.0"`
	MinPurchaseAmount  float64   `gorm:"default:0"`
}

// Acumulación de Puntos/Cashback
type Earnings struct {
	ID             uint    `gorm:"primaryKey"`
	TransactionID  uint    `gorm:"not null"`
	PointsEarned   float64 `gorm:"default:0"`
	CashbackEarned float64 `gorm:"default:0"`
}

// Recompensa
type Reward struct {
	ID             uint     `gorm:"primaryKey"`
	BusinessID     uint     `gorm:"not null"`
	Business       Business `gorm:"foreignKey:BusinessID"`
	PointsRequired int      `gorm:"not null"`
	Description    string   `gorm:"not null"`
}

// Redención de puntos/cashback
type Redemption struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
	BusinessID  uint      `gorm:"not null"`
	Business    Business  `gorm:"foreignKey:BusinessID"`
	RewardID    uint      `gorm:"not null"`
	Reward      Reward    `gorm:"foreignKey:RewardID"`
	PointsSpent int       `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
