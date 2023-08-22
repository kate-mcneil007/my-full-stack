package service

import (
	"context"

	"github.com/kate-mcneil007/my-full-stack/api-backend/pkg"
	database "github.com/kate-mcneil007/my-full-stack/api-backend/pkg/db"
)

type ServiceInterface interface {
	CreateInventoryItem() (pkg.Inventory, error)
	GetInventoryItem(context.Context, int) (pkg.Inventory, error)
	UpdateInventoryItem(context.Context, pkg.Inventory) (pkg.Inventory, error)
	DeleteInventoryItem() (bool, error)
}

type Service struct {
	db *database.DatabaseInterface
}

// NewService creates a new instance of the Service
func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateInventoryItem() (pkg.Inventory, error) {
	return pkg.Inventory{ID: 1}, nil
}

func (s *Service) GetInventoryItem(ctx context.Context, id int) (pkg.Inventory, error) {
	return pkg.Inventory{ID: id}, nil
}

func (s *Service) UpdateInventoryItem(ctx context.Context, inventory pkg.Inventory) (pkg.Inventory, error) {
	return inventory, nil
}

func (s *Service) DeleteInventoryItem() (bool, error) {
	return true, nil
}
