package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

var Conn *pgxpool.Pool
// SetupDB start DB connection
// sets Conn variable
func SetupDB() {
	os.Setenv("DATABASE_URL","postgres://lneovlgfqenrrk:57a4a43507c7bd653110519e4e70ca735c5b5263f7e766b215e1aa1637c2c086@ec2-44-199-111-161.compute-1.amazonaws.com:5432/d68joipo1e5af3")
	conn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	Conn = conn
}