package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
)

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
		ccreate table if not exists orders (
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
	return errors.New(msg)
}
