package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"pgx-pgsql/api/products"

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
	app.Use(logger.New())
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

	product := products.ProductDeps{DB: conn}
	product.ProductRouter(app)

	// #region
	// var name string
	// err = conn.QueryRow(context.Background(), "select name from products where stock = 51").Scan(&name)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Println("Creating products table", name)
	// #endregion

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})


	app.Listen(":3000")
}

//migration
func Migrate(ctx context.Context, db *pgx.Conn) error {

	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	//create users table
	_, err = tx.Exec(ctx, `
		create table if not exists users(
			id bigserial primary key,
			name varchar(255) not null,
			email varchar(255) not null,
			password text not null,
			created_at timestamp default current_timestamp,
			updated_at timestamp default current_timestamp
		);
	`)
	if err != nil {
		if errRollback := tx.Rollback(ctx); errRollback != nil {
			return errors.New("unable to rollback users table: " + errRollback.Error())
		}
		return errors.New("unable to create users table: " + err.Error())
	}

	//create products table
	_, err = tx.Exec(ctx, `
		create table if not exists products(
			id bigserial primary key,
			name varchar(255) not null,
			stock smallint not null default 0,
			price double precision not null,
			created_at timestamp default current_timestamp,
			updated_at timestamp default current_timestamp
		);
	`)
	if err != nil {
		if errRollback := tx.Rollback(ctx); errRollback != nil {
			return errors.New("unable to rollback products table: " + errRollback.Error())
		}
		return errors.New("unable to create users table: " + err.Error())
	}

	//create orders table
	_, err = tx.Exec(ctx, `
		create table if not exists orders (
			id bigserial primary key,
			user_id bigint not null,
			product_id bigint not null,
			quantity smallint not null default 0,
			total double precision not null,
			created_at timestamp default current_timestamp,
			updated_at timestamp default current_timestamp,
			CONSTRAINT fk_product
				FOREIGN KEY (product_id)
					REFERENCES products(id)
						ON DELETE CASCADE,
			CONSTRAINT fk_user
				FOREIGN KEY (user_id)
					REFERENCES users(id)
						ON DELETE CASCADE
		);
	`)
	if err != nil {
		if errRollback := tx.Rollback(ctx); errRollback != nil {
			return errors.New("unable to rollback orders table: " + errRollback.Error())
		}
		return errors.New("unable to create orders table: " + err.Error())
	}
	err = tx.Commit(ctx)
	if err != nil {
		if errRollback := tx.Rollback(ctx); errRollback != nil {
			return errors.New("unable to rollback commit transaction: " + errRollback.Error())
		}
		return errors.New("unable to commit transaction: " + err.Error())
	}

	msg := fmt.Sprintf("migration successful")
	log.Println(msg)
	return nil
}
