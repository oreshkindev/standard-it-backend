## Database

- Accepts the basic context with cancellation
- Returns a new Connection{} structure with our connection

This implementation uses PgxPool because it allows fine tuning of database behaviour and allows for multi-threading.
It doesn't matter which database driver you want to use, it can be Mongo, Arango, Maria and others.

You can create a new directory or several for different drivers inside. Each directory should contain a file with Dial() function,
and in `package database` you can switch between necessary drivers.

### Options

Here is a list of the main options:

#### MaxConns

MaxConns is the maximum number of database connections that can be opened in the connection pool at one time.
If all connections are in use, new requests will wait for a connection to be released or will be rejected if the limit is reached.
The default value is 10 times the number of CPU cores.

#### MinConns

MinConns is the minimum number of database connections that should always be open and ready to be used in the connection pool.
This value ensures that even under low load, your application will maintain the specified number of always-active connections
to ensure a fast response to incoming requests.

#### HealthCheckPeriod

HealthCheckPeriod is the time period over which the connection pool will perform a health check for each open connection.
This helps to ensure that faulty or hung connections are detected and closed, and replaced with new connections if necessary,
to keep the connection pool up and running.

#### MaxConnLifetime

MaxConnLifetime is the maximum lifetime of a connection in the pool.
When this time expires, the connection is closed and removed from the pool,
even if it is still active. This prevents potential problems associated with
long-lived connections, such as memory leaks or outdated settings.

#### MaxConnIdleTime

MaxConnIdleTime is the maximum amount of time a connection is idle after which
the connection will be closed and removed from the pool. This helps free up
resources that could have been spent on maintaining a connection that has been
idle for a long time.

#### ConnectTimeout

ConnectTimeout is the maximum time to wait for a connection to be established to the database.
If the connection cannot be established within the specified time, the connection process will be aborted with an error.

#### KeepAlive

KeepAlive specifies the time interval at which keep-alive messages will be sent to maintain the connection.
This helps to ensure that the connection is not closed by a network device or server due to inactivity.

#### Timeout

Timeout is the maximum time to wait for a connection to be established before the operation is terminated by timeout.

## Usage

```go
// Create a context that is cancellable and cancel it on exit.
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

if connection, err = database.New(ctx); err != nil {
	// There is no need to run the application without connecting to the database.
	panic(err)
}
// Close the connection when the program exits.
defer connection.Close()
```

And then:

```go
query := `
    SELECT *
    FROM users
`

// Prepare query
rows, err := connection.Query(ctx, query);
if err != nil {
	log.Println(err)
}

// Collect all records, if exist
entries, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Entry]);
if err != nil {
	log.Println(err)
}

// Print it
for _, entry := range entries {
	log.Println(entry)
}
```
