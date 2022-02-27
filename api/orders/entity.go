package orders

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4"
)

type Order struct {
	ID        int64     `json:"id" bson:"id"`
	UserID    int64     `json:"user_id" bson:"user_id"`
	ProductID int64     `json:"product_id" bson:"product_id"`
	Quantity  int64     `json:"quantity" bson:"quantity"`
	Total     float64   `json:"total" bson:"total"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type OrderRequest struct {
	UserID    int64   `json:"user_id" bson:"user_id" validate:"required"`
	ProductID int64   `json:"product_id" bson:"product_id" validate:"required"`
	Quantity  int64   `json:"quantity" bson:"quantity" validate:"required"`
	Total     float64 `json:"total" bson:"total" validate:"required"`
}

type OrderDeps struct {
	DB        *pgx.Conn
	Validator *validator.Validate
}
