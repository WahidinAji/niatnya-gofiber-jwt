package main

import (
	"github.com/jackc/pgx/v4"
)

// type Product struct {

// }

type Product struct {
	DB *pgx.Conn
}
