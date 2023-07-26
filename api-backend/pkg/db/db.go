import (
	"github.com/jackc/pgx/v5"
	"context"
	// Added this 
	"fmt" 
	"os"
)

func ConnectPostgreSqlDb(ctx context.Context) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
}
// Database operations go here...
// Can use 'conn' object to execute inserts and queries.