package service

// Business logic for controllers
// Call db file

import (
	"context"

	database "github.com/kate-mcneil007/my-full-stack/api-backend/pkg/db"
)

// HandlerInterface defines the methods that a handler should implement
type ServiceInterface interface {
	// THESE PARAMS & RESPONSES WILL BE DIFFERENT
	// Return inventory & error for now
	CreateInventoryItem() (database.Inventory, error)
	GetInventoryItem(context.Context, int) (database.Inventory, error)
	UpdateInventoryItem(context.Context, database.Inventory) (database.Inventory, error)
	DeleteInventoryItem() (bool, error)
}

// Service implements the HandlerInterface
type Service struct {
	// db db interface
	// logger ? maybe
}

// NewService creates a new instance of the Service
func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateInventoryItem() (database.Inventory, error) {
	return database.Inventory{ID: 1}, nil
}

func (s *Service) GetInventoryItem(ctx context.Context, id int) (database.Inventory, error) {
	return database.Inventory{ID: id}, nil
}

func (s *Service) UpdateInventoryItem(ctx context.Context, inventory database.Inventory) (database.Inventory, error) {
	return inventory, nil
}

func (s *Service) DeleteInventoryItem() (bool, error) {
	return true, nil
}
