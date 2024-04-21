package controllers

import "github.com/jackc/pgx/v5/pgxpool"

type AuthController struct {
	db *pgxpool.Pool
}

type AuthControllerParams struct {
	DB *pgxpool.Pool
}

type AuthControllerInterface interface {
	Login(email string, password string) error
	Register() error
}

func NewAuthController(params AuthControllerParams) AuthControllerInterface {
	return &AuthController{db: params.DB}
}

// TODO: Fix returning the user
func (c *AuthController) Login(email string, password string) error {
	return nil
}

// TODO: Fix returning the user
func (c *AuthController) Register() error {
	return nil
}
