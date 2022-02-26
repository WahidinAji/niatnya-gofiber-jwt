package orders

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)


type Repository interface {
	FindAll(ctx context.Context, db *pgx.Conn) ([]Order, error)
	FindById(ctx context.Context, db *pgx.Conn, id int64) (Order, error)
}

func FindAll(ctx context.Context, db *pgx.Conn) ([]Order, error) {
	err := db.Ping(ctx)
	if err != nil {
		return nil, errors.New("Unable to connect to database : " + err.Error())
	}

	tx, err := db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return nil, errors.New("Unable to begin transaction : " + err.Error())
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, "select * from orders")
	if err != nil {
		return nil, errors.New("Unable to query database : " + err.Error())
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.ID, &order.UserID, &order.ProductID, &order.Quantity, &order.Total, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, errors.New("Unable to scan row : " + err.Error())
		}
		orders = append(orders, order)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, errors.New("Unable to commit transaction : " + err.Error())
	}
	return orders, nil
}

func (d *OrderDeps) FindById(ctx context.Context, id int64) (Order, error) {

	var exists bool
	err := d.DB.QueryRow(ctx, "select exists(select 1 from orders where id = $1)", id).Scan(&exists)
	if err != nil {
		return Order{}, errors.New("Unable to query database : " + err.Error())
	}
	if !exists {
		return Order{}, errors.New("Order not found")
	}

	row := d.DB.QueryRow(ctx, "select * from orders where id = $1", id)
	var order Order
	err = row.Scan(&order.ID, &order.UserID, &order.ProductID, &order.Quantity, &order.Total, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return Order{}, errors.New("Unable to scan row : " + err.Error())
	}

	return order, nil
}
