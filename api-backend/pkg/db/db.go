import (
	"github.com/jackc/pgx/v5"
	"context"
)

func ConnectPostgreSqlDb(ctx context.Context) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
}