package repository

import (
	"context"
	"leal/internal/core/domain/models/db"
	"time"

	"gorm.io/gorm"
)

func (r *repository) FindBranch(ctx context.Context, branchID int) (*db.Branch, error) {
	var branch db.Branch
	err := r.db.WithContext(ctx).Where("id = ?", branchID).First(&branch).Error
	return &branch, err
}

func (r *repository) FindConversionFactor(ctx context.Context, businessTaxID int, branchID int) (*db.ConversionFactor, error) {
	var factor db.ConversionFactor
	err := r.db.WithContext(ctx).
		Where("business_tax_id = ? AND (branch_id = ? OR branch_id IS NULL)", businessTaxID, branchID).
		Order("branch_id DESC NULLS LAST, id").
		First(&factor).Error

	if err != nil {
		return nil, err
	}
	return &factor, nil
}

func (r *repository) FindActiveCampaign(ctx context.Context, branchID int, now time.Time) (*db.Campaign, error) {
	var campaign db.Campaign
	err := r.db.WithContext(ctx).Where("branch_id = ? AND start_date <= ? AND end_date >= ?", branchID, now, now).First(&campaign).Error
	if err != nil {
		return nil, err
	}
	return &campaign, err
}

func (r *repository) SaveTransaction(ctx context.Context, tx *db.Transaction, earnings *db.Earnings) error {
	return r.db.WithContext(ctx).Transaction(func(txDB *gorm.DB) error {
		if err := txDB.Create(tx).Error; err != nil {
			return err
		}
		earnings.TransactionID = tx.ID
		return txDB.Create(earnings).Error
	})
}
