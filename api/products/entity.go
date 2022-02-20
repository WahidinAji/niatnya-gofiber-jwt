package products

import (
	"time"

	"github.com/jackc/pgx/v4"
)

type Product struct {
	ID        int64     `json:"id" bson:"id"`
	Name      string    `json:"name" bson:"name"`
	Stock     int       `json:"stock" bson:"stock"`
	Price     float64   `json:"price" bson:"price"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

//DeleteRequest struct is used to parse Delete Request for Product
type DeleteRequest struct {
	ID int64 `json:"id"`
}

//Depedency struct is used to parse Dependency Request for Product
type ProductDeps struct {
	DB *pgx.Conn
}
