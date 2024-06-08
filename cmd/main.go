package main

import (
	"context"
	"standard-it-backend/external/database"
)

var (
	connection *database.Connection
	err        error
)

func main() {
	// Create a context that is cancellable and cancel it on exit.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if connection, err = database.New(ctx); err != nil {
		// There is no need to run the application without connecting to the database.
		panic(err)
	}
	// Close the connection when the program exits.
	defer connection.Close()

}
