// Spins up HTTP Server
// Go program package
package main

// Required packages for program
import (
	"context"
	"log"

	"github.com/kate-mcneil007/my-full-stack/api-backend/pkg/controller"
	database "github.com/kate-mcneil007/my-full-stack/api-backend/pkg/db"
	"github.com/kate-mcneil007/my-full-stack/api-backend/pkg/handler"
)

func main() {
	ctx := context.Background()
	conn, err := database.ConnectPostgreSqlDb(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	// handler.SetupRoutes()
	// // Once handlers set up this is called to tell global HTTP server to listen for incoming requests
	// // Here we use port :3000
	// err = http.ListenAndServe(":3000", nil)
	// // Checking for shut down/ closed server
	// // Also used to show why server stopped
	// if errors.Is(err, http.ErrServerClosed) {
	// 	fmt.Printf("server closed\n")
	// 	// Checks any other error
	// } else if err != nil {
	// 	fmt.Printf("error starting server: %s\n", err)
	// 	os.Exit(1)
	// }

	c := controller.NewController()
	handler.SetupRoutes(c)
}
