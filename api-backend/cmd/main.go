// Spins up HTTP Server
// Go program package
package main

// Required packages for program
import (
	//"context"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	database "github.com/kate-mcneil007/my-full-stack/pkg/db"
	//"github.com/jackc/pgx/v5"
)

// Both functions accept same args
// This function signature is used for HTTP handler functions
// When a request is made to the server these two vals are set up with info about the request being made
// It then calls the handler func w/ those vals
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	ctx := context.Background()
	conn, err := database.ConnectPostgreSqlDb(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	database.CreateItem(conn, &database.Item{
		ProductName: "katie is cute",
		Quantity:    666,
		Price:       9.99,
	})
	// Calls setting up handler func for request path in default server multiplexer
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	// Once handlers set up this is called to tell global HTTP server to listen for incoming requests
	// Here we use port :3000
	err = http.ListenAndServe(":3000", nil)
	// Checking for shut down/ closed server
	// Also used to show why server stopped
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
		// Checks any other error
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
