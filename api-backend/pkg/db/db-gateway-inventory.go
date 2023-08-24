package database

import (
	"context"
	"fmt"

	// pgx is how database calls are made
	"github.com/jackc/pgx/v4"
	"github.com/kate-mcneil007/my-full-stack/api-backend/pkg"
)

// ADD DB INTERFACE HERE
type DatabaseInterface interface {
	CreateInventoryItem(context.Context, pkg.Inventory) (pkg.Inventory, error)
	GetInventoryItem(context.Context, int) (pkg.Inventory, error)
	UpdateInventoryItem(context.Context, pkg.Inventory) (pkg.Inventory, error)
	DeleteInventoryItem(context.Context, int) (bool, error)
}

// Service implements the HandlerInterface
type Database struct {
	Conn *pgx.Conn
}

func NewDatabase(conn *pgx.Conn) DatabaseInterface {
	// & makes the var turn into a pointer
	return &Database{
		Conn: conn,
	}
}

func (d *Database) CreateInventoryItem(ctx context.Context, item pkg.Inventory) (pkg.Inventory, error) {
	// Return whole object
	query := "INSERT INTO inventory (product_name, quantity, price) VALUES (:product_name, :quantity, :price) RETURNING id, product_name, quanitity, price"

	// Needs to be inventory obj to be returned
	var inventory pkg.Inventory

	// Execute query using 'conn' object and the item's properties
	// Scan inventory obj
	err := d.Conn.QueryRow(context.Background(), query, item.ProductName, item.Quantity, item.Price).Scan(&inventory)
	if err != nil {
		return pkg.Inventory{}, err
	}

	return inventory, nil
}

func (d *Database) GetInventoryItem(ctx context.Context, itemID int) (pkg.Inventory, error) {
	query := "SELECT id, product_name, quantity, price FROM inventory WHERE id = :id"

	row := d.Conn.QueryRow(context.Background(), query, pkg.Inventory{ID: itemID})

	var inventory pkg.Inventory
	err := row.Scan(&inventory)
	if err != nil {
		if err == pgx.ErrNoRows {
			return pkg.Inventory{}, fmt.Errorf("item not found")
		}
		return pkg.Inventory{}, err
	}
	return inventory, nil
}

func (d *Database) UpdateInventoryItem(ctx context.Context, item pkg.Inventory) (pkg.Inventory, error) {
	query := "UPDATE inventory SET product_name = :product_name, quantity = :quantity, price = :price WHERE id = :id returning id, product_name, quanitity, price"
	var inventory pkg.Inventory
	// Execute query using 'conn' obj and the item's properties
	err := d.Conn.QueryRow(context.Background(), query, item).Scan(&inventory)
	if err != nil {
		return pkg.Inventory{}, err
	}

	return inventory, nil
}

func (d *Database) DeleteInventoryItem(ctx context.Context, itemID int) (bool, error) {
	query := "DELETE FROM inventory WHERE id = $1"

	// Execute query using 'conn' obj & item ID
	rows, err := d.Conn.Exec(context.Background(), query, itemID)
	if err != nil {
		return false, err
	}

	return rows.RowsAffected() > 0, nil
}
