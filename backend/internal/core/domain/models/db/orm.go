package db

import (
	"time"

	"gorm.io/gorm"
)

// Usuario
type User struct {
	ID             uint          `gorm:"primaryKey"`
	Name           string        `gorm:"not null"`
	DocumentNumber int           `gorm:"not null;index"`
	Email          string        `gorm:"unique;not null"`
	CreatedAt      time.Time     `gorm:"autoCreateTime"`
	Transactions   []Transaction `gorm:"foreignKey:UserID"`
	Redemptions    []Redemption  `gorm:"foreignKey:UserID"`
}

// Comercio (Business)
type Business struct {
	TaxID             int                `gorm:"primaryKey;not null"`
	Name              string             `gorm:"unique;not null"`
	Phone             int                `gorm:"not null"`
	Email             string             `gorm:"unique;not null"`
	CreatedAt         time.Time          `gorm:"autoCreateTime"`
	Branches          []Branch           `gorm:"foreignKey:BusinessTaxID;references:TaxID"`
	ConversionFactors []ConversionFactor `gorm:"foreignKey:BusinessTaxID;references:TaxID"`
	Rewards           []Reward           `gorm:"foreignKey:BusinessTaxID;references:TaxID"`
}

// Sucursal (Branch)
type Branch struct {
	ID            uint          `gorm:"primaryKey"`
	BusinessTaxID int           `gorm:"not null;index"`
	Name          string        `gorm:"not null"`
	CreatedAt     time.Time     `gorm:"autoCreateTime"`
	Business      Business      `gorm:"foreignKey:BusinessTaxID;references:TaxID"`
	Campaigns     []Campaign    `gorm:"foreignKey:BranchID"`
	Transactions  []Transaction `gorm:"foreignKey:BranchID"`
}

func (Branch) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_branch_name_business ON branches (business_tax_id, name)")
	return nil
}

// Factor de Conversión (ConversionFactor)
type ConversionFactor struct {
	ID                  uint     `gorm:"primaryKey"`
	BusinessTaxID       int      `gorm:"not null"`
	Business            Business `gorm:"foreignKey:BusinessTaxID;references:TaxID"`
	BranchID            *uint    `gorm:"index"`
	Branch              *Branch  `gorm:"foreignKey:BranchID"`
	MinAmount           float64  `gorm:"not null;default:0"`
	PointsPerCurrency   float64  `gorm:"not null"`
	CashbackPerCurrency float64  `gorm:"not null"`
}

// Compra (Transaction)
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

// Campaña (Campaign)
type Campaign struct {
	ID                 uint      `gorm:"primaryKey"`
	BusinessTaxID      int       `gorm:"not null"`
	Business           Business  `gorm:"foreignKey:BusinessTaxID;references:TaxID"`
	BranchID           uint      `gorm:"not null"`
	Branch             Branch    `gorm:"foreignKey:BranchID"`
	StartDate          time.Time `gorm:"not null"`
	EndDate            time.Time `gorm:"not null"`
	PointsMultiplier   float64   `gorm:"default:1.0"`
	CashbackMultiplier float64   `gorm:"default:1.0"`
	MinPurchaseAmount  float64   `gorm:"default:0"`
}

// Acumulación de Puntos/Cashback (Earnings)
type Earnings struct {
	ID             uint    `gorm:"primaryKey"`
	TransactionID  uint    `gorm:"not null"`
	PointsEarned   float64 `gorm:"default:0"`
	CashbackEarned float64 `gorm:"default:0"`
}

// Recompensa (Reward)
type Reward struct {
	ID             uint     `gorm:"primaryKey"`
	BusinessTaxID  string   `gorm:"not null"`
	Business       Business `gorm:"foreignKey:BusinessTaxID;references:TaxID"`
	PointsRequired int      `gorm:"not null"`
	Description    string   `gorm:"not null"`
}

// Redención de puntos/cashback (Redemption)
type Redemption struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `gorm:"not null"`
	User          User      `gorm:"foreignKey:UserID"`
	BusinessTaxID string    `gorm:"not null"`
	Business      Business  `gorm:"foreignKey:BusinessTaxID;references:TaxID"`
	RewardID      uint      `gorm:"not null"`
	Reward        Reward    `gorm:"foreignKey:RewardID"`
	PointsSpent   int       `gorm:"not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}
