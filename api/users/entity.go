package users

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4"
)

type User struct {
	ID        int64     `json:"id" bson:"id"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type LoginInput struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=8"`
}

type UserResponse struct {
	Email string `json:"email" bson:"email"`
	Name  string `json:"name" bson:"name"`
}

type UserDeps struct {
	DB        *pgx.Conn
	Validator *validator.Validate
}
