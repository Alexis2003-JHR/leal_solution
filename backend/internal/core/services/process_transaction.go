package services

import (
	"context"
	"fmt"
	"leal/internal/core/domain/models/request"
)

func (s *service) ProcessTransaction(ctx context.Context, request request.ProcessTransaction) error {
	branch, err := s.repo.FindBranch(ctx, request.BranchID)
	if err != nil {
		return err
	}

	fmt.Println(branch)

	return nil
}
