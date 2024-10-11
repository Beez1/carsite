package cloudsql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

// ConnectWithConnector establishes a connection to the PostgreSQL instance using
// the Cloud SQL Go Connector without requiring the Cloud SQL Proxy.
func ConnectWithConnector() (*sql.DB, error) {
	// Helper function to retrieve environment variables
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error: %s environment variable not set.\n", k)
		}
		return v
	}

	// Load environment variables
	var (
		dbUser                 = mustGetenv("DB_USER")                  // e.g., 'beez'
		dbPwd                  = mustGetenv("DB_PASS")                  // e.g., 'carsitetest1234'
		dbName                 = mustGetenv("DB_NAME")                  // e.g., 'carsite'
		instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME") // e.g., 'dotted-chiller-437712-q3:europe-west2:carsite'
		usePrivate             = os.Getenv("PRIVATE_IP")                // Optional: Set to use private IP
	)

	// Create PostgreSQL connection string
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s", dbUser, dbPwd, dbName)
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("Error parsing DSN: %v", err)
	}

	// Set up Cloud SQL Connector options
	var opts []cloudsqlconn.Option
	if usePrivate != "" {
		opts = append(opts, cloudsqlconn.WithDefaultDialOptions(cloudsqlconn.WithPrivateIP()))
	}

	// Create a new Cloud SQL Connector dialer
	d, err := cloudsqlconn.NewDialer(context.Background(), opts...)
	if err != nil {
		return nil, fmt.Errorf("Error creating Cloud SQL Dialer: %v", err)
	}

	// Use Cloud SQL Connector DialFunc to connect to the instance
	config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
		return d.Dial(ctx, instanceConnectionName)
	}

	// Register the pgx connection with the stdlib
	dbURI := stdlib.RegisterConnConfig(config)
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("Error opening SQL connection: %w", err)
	}

	// Test the connection
	if err := dbPool.Ping(); err != nil {
		return nil, fmt.Errorf("Error pinging database: %v", err)
	}

	log.Println("Successfully connected to the database!")
	return dbPool, nil
}


