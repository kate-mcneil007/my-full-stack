// Spins up HTTP Server
// Go program package
package main

// Required packages for program
import (
	"context"
	"log"

	controller "github.com/kate-mcneil007/my-full-stack/api-backend/pkg/controller"
	database "github.com/kate-mcneil007/my-full-stack/api-backend/pkg/db"
	"github.com/kate-mcneil007/my-full-stack/api-backend/pkg/handler"
	"github.com/kate-mcneil007/my-full-stack/api-backend/pkg/service"
)

func main() {
	// returns a non-nil, empty Context
	// Is never canceled, has no values, and has no deadline
	// Typically used as the top-level Context for incoming requests
	ctx := context.Background()

	// Variable named conn to store the connection to the PostgreSQL database
	// conn holds connection object used to interact with the database
	// err stores potential errors
	// Our non-nil, empty context variable is passed into func ConnectPostgresSqlDb found in package "database"
	conn, err := database.ConnectPostgreSqlDb(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// NewService establishes new service structure
	// This is then passed to create a new instance of the Controller
	// Routes are then set up using this new controller isntance
	d := database.NewDatabase(conn)
	s := service.NewService(d)
	c := controller.NewController(s)
	handler.SetupRoutes(c)
}
