package users

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v4"
)

type JwtUserClaim struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	jwt.StandardClaims
}

type JwtResponse struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Token string `json:"token"`
}

type ValidationResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Error string `json:"error"`
	Value string `json:"value"`
}

func (d *UserDeps) ValidateStruct(user LoginInput) []*ValidationResponse {
	var errors []*ValidationResponse
	err := d.Validator.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationResponse
			// element.FailedField = err.StructNamespace()
			element.Field = err.Field()
			element.Tag = err.Tag()
			element.Error = err.Error()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func (d *UserDeps) UValidate(user LoginInput) string {
	var errs string
	err := d.Validator.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errs = errs + err.Field() + ", " + err.Tag() + ", " + err.Param() + ", " + err.Error() + "; "
			// errs = append(errs, err.Field()+", "+err.Tag()+", "+err.Param()+", "+err.Error())
		}
		return errs
	}
	return ""
}

func (d *UserDeps) LoginServiceUser(ctx context.Context, input LoginInput) (*JwtResponse, error) {
	err := d.DB.Ping(ctx)
	if err != nil {
		return nil, errors.New("Unable to connect to database : " + err.Error())
	}

	// errs := d.UValidate(LoginInput{Email: email, Password: password})
	errs := d.UValidate(input)
	fmt.Println(errs)
	if errs != "" {
		return nil, fiber.NewErrors(fiber.StatusBadRequest, errs)
	}

	tx, err := d.DB.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return nil, errors.New("Unable to begin transaction : " + err.Error())
	}
	defer tx.Rollback(ctx)

	var checkPass bool
	pass, err := d.GetPassUser(ctx, input.Email)

	checkPass = CheckPassword(input.Password, pass)
	if !checkPass {
		return nil, errors.New("Invalid email or password")
	}

	_, err = d.CheckRepoUser(ctx, input.Email, string(pass))
	if err != nil {
		return nil, err
	}

	user, err := d.LoginUserRepo(ctx, input.Email, string(pass))
	if err != nil {
		return nil, err
	}

	claims := &JwtUserClaim{
		user.Name,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, errors.New("Unable to sign token : " + err.Error())
	}
	var response JwtResponse
	response.Name = user.Name
	response.Email = user.Email
	response.Token = tokenString
	return &response, nil
}
