package controllers

import (
	"context"
	"errors"
	"fmt"

	"github.com/MalshanPerera/go-expense-tracker/database/sqlc"
	"github.com/MalshanPerera/go-expense-tracker/modals"
	"github.com/MalshanPerera/go-expense-tracker/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserResponse modals.UserResponse

type AuthController struct {
	db      *pgxpool.Pool
	queries *sqlc.Queries
}

type AuthControllerParams struct {
	DB      *pgxpool.Pool
	Queries *sqlc.Queries
}

type AuthControllerInterface interface {
	Login(ctx context.Context, userPayload modals.LoginUserParams) (*UserResponse, error)
	Register(ctx context.Context, userPayload modals.CreateUserParams) (*UserResponse, error)
}

func NewAuthController(params AuthControllerParams) AuthControllerInterface {
	return &AuthController{db: params.DB, queries: params.Queries}
}

func (c *AuthController) Login(ctx context.Context, userPayload modals.LoginUserParams) (*UserResponse, error) {

	user, err := c.queries.GetUserByEmail(
		ctx,
		userPayload.Email,
	)

	if err != nil {
		if user.ID.Bytes == [16]byte{} {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if !utils.ComparePasswords(user.Password, []byte(userPayload.Password)) {
		return nil, errors.New("invalid email or password")
	}

	accessToken, expiredAt, err := utils.CreateToken(user.ID, utils.Type.AccessToken)
	if err != nil {
		return nil, err
	}

	refreshToken, _, err := utils.CreateToken(user.ID, utils.Type.RefreshToken)
	if err != nil {
		return nil, err
	}

	session := sqlc.UpdateSessionParams{
		UserID:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiredAt,
	}

	newSession, err := c.queries.UpdateSession(
		ctx,
		session,
	)
	if err != nil {
		return nil, err
	}

	loggedInUser := &UserResponse{
		ID:           user.ID,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		AccessToken:  newSession.AccessToken,
		RefreshToken: newSession.RefreshToken,
		ExpiredAt:    newSession.ExpiresAt,
		CreatedAt:    user.CreatedAt,
	}

	return loggedInUser, nil
}

func (c *AuthController) Register(ctx context.Context, userPayload modals.CreateUserParams) (*UserResponse, error) {

	_, err := c.queries.GetUserByEmail(ctx, userPayload.Email)
	if err == nil {
		return nil, fmt.Errorf("user with email %s already exists", userPayload.Email)
	}

	hashedPassword, err := utils.HashPassword(userPayload.Password)
	if err != nil {
		return nil, err
	}

	user := sqlc.CreateUserParams{
		FirstName: userPayload.FirstName,
		LastName:  userPayload.LastName,
		Email:     userPayload.Email,
		Password:  hashedPassword,
	}

	createdUser, err := c.queries.CreateUser(
		ctx,
		user,
	)
	if err != nil {
		return nil, err
	}

	accessToken, expiredAt, err := utils.CreateToken(createdUser.ID, utils.Type.AccessToken)
	if err != nil {
		return nil, err
	}

	refreshToken, _, err := utils.CreateToken(createdUser.ID, utils.Type.RefreshToken)
	if err != nil {
		return nil, err
	}

	session := sqlc.CreateSessionParams{
		UserID:       createdUser.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiredAt,
	}

	newSession, err := c.queries.CreateSession(
		ctx,
		session,
	)
	if err != nil {
		return nil, err
	}

	registeredUser := &UserResponse{
		ID:           createdUser.ID,
		Email:        createdUser.Email,
		FirstName:    createdUser.FirstName,
		LastName:     createdUser.LastName,
		AccessToken:  newSession.AccessToken,
		RefreshToken: newSession.RefreshToken,
		ExpiredAt:    newSession.ExpiresAt,
		CreatedAt:    createdUser.CreatedAt,
	}

	return registeredUser, nil
}
