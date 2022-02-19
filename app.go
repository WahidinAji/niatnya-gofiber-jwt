package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v4"
)

func init() {
	if os.Getenv("DB_URL") == "" {
		log.Fatal("DB_URL is not set")
	}
}

func main() {
	app := fiber.New()

	ctx := context.Background()

	db := os.Getenv("DB_URL")
	conn, err := pgx.Connect(ctx, db)

	err = Migrate(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(ctx)

	// var name string

	// err = conn.QueryRow(context.Background(), "select name from products where stock = 51").Scan(&name)

	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println("Creating products table", name)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use(logger.New())

	app.Listen(":3000")
}
