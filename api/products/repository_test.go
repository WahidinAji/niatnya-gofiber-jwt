package products

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v4"
)

//create test function to get all products with product dependencies
func TestFindAll(t *testing.T) {
	ctx := context.Background()
	db := "postgres://postgres:postgres@localhost:5432/fiber_jwt?sslmode=disable"
	conn, err := pgx.Connect(ctx, db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	product := ProductDeps{DB: conn}
	rows, err := product.FindAll(ctx)
	if err != nil {
		t.Errorf("Unable to query database : %v", err)
	}
	if len(rows) > 0 {
		t.Errorf("Expected nil row, got %v", len(rows))
	}

	for _, row := range rows {
		t.Logf("%+v", row)
	}

	conn.Exec(ctx, "insert into products (name, stock, price) values ($1, $2, $3)", "test", 51, 1.99)
	rows, err = product.FindAll(ctx)
	if err != nil {
		t.Errorf("Unable to query database : %v", err)
	}
	if len(rows) == 0 {
		t.Errorf("Expected at least one row, got %v", len(rows))
	}

	conn.Exec(ctx, "delete from products")
	rows, err = product.FindAll(ctx)
	if err != nil {
		t.Errorf("Unable to query database : %v", err)
	}
	if len(rows) > 0 {
		t.Errorf("Expected nil row row, got %v", len(rows))
	}
}
