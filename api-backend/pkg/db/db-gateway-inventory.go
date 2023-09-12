package database

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/kate-mcneil007/my-full-stack/api-backend/pkg"
)

type DatabaseInterface interface {
	CreateInventoryItem(context.Context, pkg.Inventory) (pkg.Inventory, error)
	GetInventoryItem(context.Context, int) (pkg.Inventory, error)
	UpdateInventoryItem(context.Context, pkg.Inventory) (pkg.Inventory, error)
	DeleteInventoryItem(context.Context, int) (bool, error)
}

type Database struct {
	Conn *sqlx.DB
}

func NewDatabase(conn *sqlx.DB) *Database {
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
	rows, err := d.Conn.NamedQueryContext(ctx, query, item)
	if err != nil {
		return pkg.Inventory{}, err
	}
	err = rows.StructScan(&inventory)
	if err != nil {
		return pkg.Inventory{}, err
	}

	return inventory, nil
}

func (d *Database) GetInventoryItem(ctx context.Context, itemID int) (pkg.Inventory, error) {
	query := "SELECT * FROM inventory WHERE id = :id"

	var inventory []pkg.Inventory
	rows, err := d.Conn.NamedQueryContext(ctx, query, pkg.Inventory{ID: itemID})
	if err != nil {
		log.Fatal("i am here " + err.Error())
		return pkg.Inventory{}, err
	}
	defer rows.Close()
	for rows.Next() {
		record := pkg.Inventory{}
		err := rows.StructScan(&record)
		if err != nil {
			return pkg.Inventory{}, err
		}
		inventory = append(inventory, record)
	}

	return inventory[0], nil
}

func (d *Database) UpdateInventoryItem(ctx context.Context, item pkg.Inventory) (pkg.Inventory, error) {
	query := "UPDATE inventory SET product_name = :product_name, quantity = :quantity, price = :price WHERE id = :id returning id, product_name, quanitity, price"
	var inventory pkg.Inventory
	// Execute query using 'conn' obj and the item's properties
	rows, err := d.Conn.NamedQueryContext(ctx, query, item)
	if err != nil {
		return pkg.Inventory{}, err
	}
	err = rows.StructScan(&inventory)
	if err != nil {
		return pkg.Inventory{}, err
	}

	return inventory, nil
}

func (d *Database) DeleteInventoryItem(ctx context.Context, itemID int) (bool, error) {
	query := "DELETE FROM inventory WHERE id = $1"

	// Execute query using 'conn' obj & item ID
	result, err := d.Conn.NamedExecContext(context.Background(), query, itemID)
	if err != nil {
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
