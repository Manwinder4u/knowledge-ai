package database

import "github.com/jackc/pgx/v4/pgxpool"

type DB struct {
	Conn *pgxpool.Pool
}
