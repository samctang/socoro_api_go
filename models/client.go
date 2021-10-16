package models

import "github.com/jmoiron/sqlx"

type Client struct {
	DB *sqlx.DB
}

// NewClient is function to create new instance.
func NewClient(db *sqlx.DB) (c Client) {
	return Client{
		DB: db,
	}
}
