package orders

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4"
)

func (d *OrderDeps) OrderValidation(order OrderRequest) string {
	var errs string
	err := d.Validator.Struct(order)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errs = errs + err.Field() + ", " + err.Tag() + ", " + err.Param() + ", " + err.Error() + "; "
			// errs = append(errs, err.Field()+", "+err.Tag()+", "+err.Param()+", "+err.Error())
		}
		return errs
	}
	return ""
}

func (d *OrderDeps) ServiceById(ctx context.Context, id int64) (Order, error) {
	err := d.DB.Ping(ctx)
	if err != nil {
		return Order{}, errors.New("Unable to connect to database : " + err.Error())
	}

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

func (d *OrderDeps) ServiceCreateOrder(ctx context.Context, orderRequest OrderRequest) (OrderRequest, error) {

	err := d.DB.Ping(ctx)
	if err != nil {
		return OrderRequest{}, errors.New("Unable to connect to database : " + err.Error())
	}

	errs := d.OrderValidation(orderRequest)
	if errs != "" {
		return OrderRequest{}, errors.New("validation error: " + errs)
	}

	tx, err := d.DB.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return OrderRequest{}, errors.New("Unable to begin transaction : " + err.Error())
	}
	defer tx.Rollback(ctx)

	order, err := d.RepoCreateOrder(ctx, orderRequest)
	if err != nil {
		return OrderRequest{}, errors.New("Unable to query database : " + err.Error())
	}

	err = tx.Commit(ctx)
	if err != nil {
		return OrderRequest{}, errors.New("Unable to commit transaction : " + err.Error())
	}

	return order, nil
}
