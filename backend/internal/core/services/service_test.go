package services_test

import (
	"context"
	"testing"
	"time"

	"leal/internal/core/domain/models/db"
	"leal/internal/core/domain/models/request"
	"leal/internal/core/domain/ports/mocks"
	"leal/internal/core/services"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateBranch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockRepository(ctrl)

	repoMock.EXPECT().
		InsertBranch(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, branch *db.Branch) error {
			branch.ID = 1
			return nil
		})

	repoMock.EXPECT().
		InsertConversionFactor(gomock.Any(), gomock.Any()).
		Return(nil)

	service := services.NewService(repoMock)

	req := request.CreateBranch{
		NITEmpresa:     12345678,
		NombreSucursal: "Sucursal Central",
		ConversionFactor: request.ConversionFactor{
			MinAmount:           100,
			PointsPerCurrency:   10,
			CashbackPerCurrency: 0.05,
		},
	}
	err := service.CreateBranch(context.Background(), req)

	assert.NoError(t, err)
}

func TestCreateBusiness(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockRepository(ctrl)

	service := services.NewService(repoMock)

	req := request.CreateBusiness{
		RazonSocial: "Empresa XYZ",
		NIT:         987654321,
		Telefono:    123456789,
		Correo:      "contacto@empresa.xyz",
		ConversionFactor: request.ConversionFactor{
			MinAmount:           200,
			PointsPerCurrency:   20,
			CashbackPerCurrency: 0.10,
		},
	}

	repoMock.EXPECT().InsertBusiness(gomock.Any(), gomock.Any()).Return(nil).Times(1)
	repoMock.EXPECT().InsertConversionFactor(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	err := service.CreateBusiness(context.Background(), req)

	assert.NoError(t, err)
}

func TestCreateCampaign(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockRepository(ctrl)

	service := services.NewService(repoMock)

	req := request.CreateCampaign{
		TaxID:              987654321,
		BranchID:           1,
		StartDate:          "2025-03-21",
		EndDate:            "2025-04-21",
		PointsMultiplier:   1.5,
		CashbackMultiplier: 0.05,
		MinPurchaseAmount:  500,
	}

	repoMock.EXPECT().InsertCampaign(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	err := service.CreateCampaign(context.Background(), req)

	assert.NoError(t, err)
}

func TestCreateReward(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockRepository(ctrl)

	service := services.NewService(repoMock)

	req := request.CreateReward{
		Name:           "Recompensa 1",
		Description:    "Descripción de la recompensa",
		PointsRequired: 100,
		BusinessTaxID:  123456789,
	}

	repoMock.EXPECT().InsertReward(gomock.Any(), gomock.Any()).Return(nil)

	err := service.CreateReward(context.Background(), req)

	assert.NoError(t, err)
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockRepository(ctrl)

	service := services.NewService(repoMock)

	req := request.CreateUser{
		Name:           "Juan Pérez",
		DocumentNumber: 123456789,
		Email:          "juan.perez@example.com",
	}

	repoMock.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(nil)

	err := service.CreateUser(context.Background(), req)

	assert.NoError(t, err)
}

func TestObtainBranches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockRepository(ctrl)
	service := services.NewService(repoMock)

	taxID := 123456789

	dbBranches := []db.Branch{
		{ID: 1, Name: "Sucursal 1"},
		{ID: 2, Name: "Sucursal 2"},
	}

	repoMock.EXPECT().GetBranches(gomock.Any(), taxID).Return(dbBranches, nil)

	branches, err := service.ObtainBranches(context.Background(), taxID)

	assert.NoError(t, err)
	assert.Len(t, branches, 2)
	assert.Equal(t, "Sucursal 1", branches[0].Nombre)
}

func TestRedeemPoints(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockRepository(ctrl)
	service := services.NewService(repoMock)

	req := request.RedeemPoints{
		User: request.CreateUser{
			DocumentNumber: 123456789,
		},
		RewardID:      1,
		BusinessTaxID: 123,
	}

	userBalance := db.UserBalance{
		Points: 100,
	}

	reward := db.Reward{
		ID:             1,
		PointsRequired: 50,
	}

	repoMock.EXPECT().GetUserBalance(gomock.Any(), req.User.DocumentNumber).Return(userBalance, nil)
	repoMock.EXPECT().GetReward(gomock.Any(), req.RewardID).Return(reward, nil)
	repoMock.EXPECT().UpdateUserRedeemPoints(gomock.Any(), req.User.DocumentNumber, reward.PointsRequired).Return(nil)
	repoMock.EXPECT().InsertRedemption(gomock.Any(), gomock.Any()).Return(nil)

	pointsSpent, err := service.RedeemPoints(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, reward.PointsRequired, pointsSpent)
}

func TestObtainCampaign(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockRepository(ctrl)
	service := services.NewService(repoMock)

	taxID := 123
	var branchIDUint uint = 10
	branchIDInt := int(branchIDUint)

	campaignsDB := []db.Campaign{
		{
			ID:                 1,
			BusinessTaxID:      taxID,
			BranchID:           &branchIDUint, // Pasamos la referencia como *uint
			StartDate:          time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC),
			EndDate:            time.Date(2025, 3, 31, 23, 59, 59, 0, time.UTC),
			PointsMultiplier:   1.5,
			CashbackMultiplier: 2.0,
			MinPurchaseAmount:  100.0,
		},
	}

	repoMock.EXPECT().GetCampaigns(gomock.Any(), taxID, gomock.Any()).Return(campaignsDB, nil)

	campaigns, err := service.ObtainCampaign(context.Background(), taxID, &branchIDInt)

	assert.NoError(t, err)
	assert.Len(t, campaigns, len(campaignsDB))

	assert.Equal(t, campaignsDB[0].ID, campaigns[0].ID)

	if campaignsDB[0].BranchID != nil {
		assert.NotNil(t, campaigns[0].BranchID)
	} else {
		assert.Nil(t, campaigns[0].BranchID)
	}

	assert.Equal(t, campaignsDB[0].StartDate.Format("2006-01-02"), campaigns[0].StartDate)
	assert.Equal(t, campaignsDB[0].EndDate.Format("2006-01-02"), campaigns[0].EndDate)
	assert.Equal(t, campaignsDB[0].PointsMultiplier, campaigns[0].PointsMultiplier)
	assert.Equal(t, campaignsDB[0].CashbackMultiplier, campaigns[0].CashbackMultiplier)
	assert.Equal(t, campaignsDB[0].MinPurchaseAmount, campaigns[0].MinPurchaseAmount)
}
