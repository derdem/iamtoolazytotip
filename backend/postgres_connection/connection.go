package postgres_connection

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

func GetConnection() *pgx.Conn {
	// Connect to the database
	config := pgx.ConnConfig{
		Host:     "database",
		Port:     5432,
		User:     "iamtoolazytotip",
		Password: "iamtoolazytotip",
		Database: "tournaments",
	}
	conn, err := pgx.Connect(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return conn
}

func GetConnectionForTest() *pgx.Conn {
	// Connect to the database from localhost
	config := pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "iamtoolazytotip",
		Password: "iamtoolazytotip",
		Database: "tournaments",
	}
	conn, err := pgx.Connect(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return conn
}
