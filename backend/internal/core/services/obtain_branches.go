package services

import (
	"context"
	"leal/internal/core/domain/models"
)

func (s *service) ObtainBranches(ctx context.Context, taxID int) ([]models.Branch, error) {
	var branches []models.Branch
	dbBranches, err := s.repo.GetBranches(ctx, taxID)
	if err != nil {
		return []models.Branch{}, err
	}

	for _, dbBranch := range dbBranches {
		branch := models.Branch{
			ID:     int(dbBranch.ID),
			Nombre: dbBranch.Name,
		}
		branches = append(branches, branch)
	}

	return branches, nil
}
