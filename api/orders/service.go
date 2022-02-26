package orders

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func (d *OrderDeps) ServiceById(ctx context.Context, id int64) (Order, error) {
	tx, err := d.DB.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return Order{}, errors.New("Unable to begin transaction : " + err.Error())
	}
	defer tx.Rollback(ctx)

	order, err := d.FindById(ctx, id)
	if err != nil {
		return Order{}, errors.New("Unable to query database : " + err.Error())
	}
	err = tx.Commit(ctx)
	if err != nil {
		return Order{}, errors.New("Unable to commit transaction : " + err.Error())
	}

	return order, nil
}
