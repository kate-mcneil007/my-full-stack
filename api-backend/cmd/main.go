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
	ctx := context.Background()
	conn, err := database.ConnectPostgreSqlDb(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	s := service.NewService()
	c := controller.NewController(s)
	handler.SetupRoutes(c)
}
