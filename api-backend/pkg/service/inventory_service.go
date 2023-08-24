package service

import (
	"context"

	"github.com/kate-mcneil007/my-full-stack/api-backend/pkg"
	database "github.com/kate-mcneil007/my-full-stack/api-backend/pkg/db"
)

type ServiceInterface interface {
	CreateInventoryItem(context.Context, pkg.Inventory) (pkg.Inventory, error)
	GetInventoryItem(context.Context, int) (pkg.Inventory, error)
	UpdateInventoryItem(context.Context, pkg.Inventory) (pkg.Inventory, error)
	DeleteInventoryItem(context.Context, int) (bool, error)
}

type Service struct {
	db *database.DatabaseInterface
}

// NewService creates a new instance of the Service
func NewService(db database.DatabaseInterface) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateInventoryItem(ctx context.Context, inventory pkg.Inventory) (pkg.Inventory, error) {
	return s.db.CreateInventoryItem(ctx, inventory)
}

func (s *Service) GetInventoryItem(ctx context.Context, id int) (pkg.Inventory, error) {
	return s.db.GetInventoryItem(ctx, id)
}

func (s *Service) UpdateInventoryItem(ctx context.Context, inventory pkg.Inventory) (pkg.Inventory, error) {
	return s.db.UpdateInventoryItem(ctx, inventory)
}

func (s *Service) DeleteInventoryItem(ctx context.Context, id int) (bool, error) {
	return s.db.DeleteInventoryItem(ctx, id)
}
