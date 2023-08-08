// Is this right?
package database

import (
	"context"

	"github.com/jackc/pgx/v4"

	// Added this
	"fmt"
	"os"
)

//	func ConnectPostgreSqlDb(ctx context.Context) {
//		conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
//		if err != nil {
//			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
//			os.Exit(1)
//		}
//		defer conn.Close(ctx)
//	}
func ConnectPostgreSqlDb(ctx context.Context) (*pgx.Conn, error) {
	url := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	return conn, nil
}

// Can use 'conn' object to execute inserts and queries.

// Entry in inventory table
type Item struct {
	ID          int     `json:"id" db:"id"`
	ProductName string  `json:"product_name" db:"product_name"`
	Quantity    int     `json:"quantity" db:"quantity"`
	Price       float64 `json:"price" db:"price"`
}

func CreateItem(conn *pgx.Conn, item *Item) error {
	query := "INSERT INTO inventory (product_name, quantity, price) VALUES ($1, $2, $3)"

	// Execute query using 'conn' object and the item's properties
	_, err := conn.Exec(context.Background(), query, item.ProductName, item.Quantity, item.Price)
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(conn *pgx.Conn) ([]Item, error) {
	var items []Item

	query := "SELECT id, product_name, quantity, price FROM inventory"

	// Execute query using 'conn' object
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result rows and populate the items slice
	for rows.Next() {
		var item Item
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

func UpdateItem(conn *pgx.Conn, item *Item) error {
	query := "UPDATE inventory SET product_name = $1, quantity = $2, price = $3 WHERE id = $4"

	// Execute query using 'conn' obj and the item's properties
	_, err := conn.Exec(context.Background(), query, item.ProductName, item.Quantity, item.Price, item.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteItem(conn *pgx.Conn, itemID int) error {
	query := "DELETE FROM inventory WHERE id = $1"

	// Execute query using 'conn' obj & item ID
	_, err := conn.Exec(context.Background(), query, itemID)
	if err != nil {
		return err
	}

	return nil
}
