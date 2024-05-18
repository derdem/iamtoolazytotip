package postgres_connection

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

func getConfig(host string) pgx.ConnConfig {
	config := pgx.ConnConfig{
		Host:     host,
		Port:     5432,
		User:     "iamtoolazytotip",
		Password: "iamtoolazytotip",
		Database: "tournaments",
	}
	return config
}

func GetConnection() *pgx.Conn {
	// Connect to the database
	config := getConfig("database")
	conn, err := pgx.Connect(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return conn
}

func GetConnectionForTest() *pgx.Conn {
	// Connect to the database from localhost
	config := getConfig("localhost")
	conn, err := pgx.Connect(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return conn
}
