package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"standard-it-backend/external/database"
	"standard-it-backend/external/router"
)

var (
	connection *database.Connection
	mux        *router.Mux

	err error
)

func main() {
	// Create a context that is cancellable and cancel it on exit.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Run the application
	if err = run(ctx); err != nil {
		log.Println(err)
	}
}

func run(ctx context.Context) error {
	// Create a new database connection.
	if connection, err = database.New(ctx); err != nil {
		// There is no need to run the application without connecting to the database.
		panic(err)
	}
	// Close the connection when the program exits.
	defer connection.Close()

	// Create a new router with the necessary middlewares and routes.
	if mux, err = router.New(ctx); err != nil {
		log.Println(err)
	}

	return http.ListenAndServe(os.Getenv("SERVICE_PORT"), mux)
}
