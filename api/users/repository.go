package users

import (
	"context"
	"errors"
)

func (d *UserDeps) CheckRepoUser(ctx context.Context, email, password string) (bool, error) {
	var authenticated bool
	err := d.DB.QueryRow(ctx, "select exists(select 1 from users where email = $1 and password = $2)", email, password).Scan(&authenticated)
	if err != nil {
		return false, errors.New("Unable to query database : " + err.Error())
	}
	if !authenticated {
		return false, errors.New("User not found")
	}
	return true, nil
}

func (d *UserDeps) LoginUserRepo(ctx context.Context, email, password string) (*UserResponse, error) {
	var user UserResponse
	err := d.DB.QueryRow(ctx, "select name, email from users where email = $1 and password = $2", email, password).Scan(&user.Email, &user.Name)
	if err != nil {
		return nil, errors.New("Unable to query database : " + err.Error())
	}
	return &user, nil
}
