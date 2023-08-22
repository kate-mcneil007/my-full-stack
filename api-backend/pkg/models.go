package pkg

// ANYWHERE WHERE THIS STRUCT IS USED NOW NEEDS TO IMPORT PKG
type Inventory struct {
	ID          int     `json:"id" db:"id"`
	ProductName string  `json:"product_name" db:"product_name"`
	Quantity    int     `json:"quantity" db:"quantity"`
	Price       float64 `json:"price" db:"price"`
}
