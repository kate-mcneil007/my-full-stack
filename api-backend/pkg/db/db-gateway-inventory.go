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
	GetItemByID(conn *pgx.Conn, itemID int) (*pkg.Inventory, error)
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
	query := "INSERT INTO inventory (product_name, quantity, price) VALUES ($1, $2, $3)"

	// Execute query using 'conn' object and the item's properties
	_, err := d.Conn.Exec(context.Background(), query, item.ProductName, item.Quantity, item.Price)
	if err != nil {
		return err
	}

	return
}

func GetInventoryItem(ctx context.Context, item int) (pkg.Inventory, error) {
	var items []pkg.Inventory

	query := "SELECT id, product_name, quantity, price FROM inventory"

	// Execute query using 'conn' object
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result rows and populate the items slice
	for rows.Next() {
		var item pkg.Inventory
		err := rows.Scan(&item.ID, &item.ProductName, &item.Quantity, &item.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func UpdateInventoryItem(context.Context, pkg.Inventory) (pkg.Inventory, error) {
	query := "UPDATE inventory SET product_name = $1, quantity = $2, price = $3 WHERE id = $4"

	// Execute query using 'conn' obj and the item's properties
	_, err := conn.Exec(context.Background(), query, item.ProductName, item.Quantity, item.Price, item.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteInventoryItem(context.Context, int) (bool, error) {
	query := "DELETE FROM inventory WHERE id = $1"

	// Execute query using 'conn' obj & item ID
	_, err := conn.Exec(context.Background(), query, itemID)
	if err != nil {
		return err
	}

	return nil
}

func GetItemByID(conn *pgx.Conn, itemID int) (*pkg.Inventory, error) {
	query := "SELECT id, product_name, quantity, price FROM inventory WHERE id = $1"

	row := conn.QueryRow(context.Background(), query, itemID)

	var inventory pkg.Inventory
	err := row.Scan(&itemID, &inventory.ProductName, &inventory.Quantity, &inventory.Price)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("item not found")
		}
		return nil, err
	}

	return &inventory, nil
}
