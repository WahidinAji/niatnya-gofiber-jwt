package products

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

//create function to Find all products with product dependencies
func (d *ProductDeps) FindAll(ctx context.Context) ([]Product, error) {
	err := d.DB.Ping(ctx)
	if err != nil {
		return nil, errors.New("Unable to connect to database : " + err.Error())
	}

	tx, err := d.DB.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return nil, errors.New("Unable to begin transaction : " + err.Error())
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, "select * from products")
	if err != nil {
		return nil, errors.New("Unable to query database : " + err.Error())
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Stock, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, errors.New("Unable to scan row : " + err.Error())
		}
		products = append(products, product)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, errors.New("Unable to commit transaction : " + err.Error())
	}
	return products, nil
}
