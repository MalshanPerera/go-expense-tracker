package controllers

import (
	"errors"
	"fmt"

	"github.com/MalshanPerera/go-expense-tracker/database/schema"
	"gorm.io/gorm"
)

type AuthController struct {
	db *gorm.DB
}

type AuthControllerParams struct {
	DB *gorm.DB
}

type AuthControllerInterface interface {
	Login(email string, password string) (*schema.User, error)
	Register(user schema.User) (*schema.User, error)
}

// TODO: Fix returning the user
func NewAuthController(params AuthControllerParams) AuthControllerInterface {
	return &AuthController{db: params.DB}
}

func (c *AuthController) Login(email string, password string) (*schema.User, error) {

	var user schema.User
	c.db.Where("email = ? AND password = ?", email, password).Find(&user)

	if user.ID != 0 {
		return nil, errors.New("invalid email or password")
	}

	return &user, nil
}

// TODO: Fix returning the user
func (c *AuthController) Register(user schema.User) (*schema.User, error) {

	result := c.db.Where("email = ?", user.Email).First(&schema.User{})

	if result.Error != nil {
		createResult := c.db.Create(&user)
		if createResult.Error != nil {
			return nil, errors.New("failed to create the user")
		}
		return &schema.User{}, nil
	}

	return nil, fmt.Errorf("user with email %s already exists", user.Email)
}
