package repository

import (
	"context"
	"leal/internal/core/domain/models/db"
)

func (r *repository) FindBranch(ctx context.Context, branchID int) (*db.Branch, error) {
	var branch db.Branch
	err := r.db.WithContext(ctx).Where("id = ?", branchID).First(&branch).Error
	return &branch, err
}
