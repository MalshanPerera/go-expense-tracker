package modals

import "github.com/jackc/pgx/v5/pgtype"

type UserResponse struct {
	ID           pgtype.UUID      `json:"id"`
	Email        string           `json:"email"`
	FirstName    string           `json:"first_name"`
	LastName     string           `json:"last_name"`
	AccessToken  string           `json:"access_token"`
	RefreshToken string           `json:"refresh_token"`
	ExpiredAt    int64            `json:"expired_at"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
}

type LoginUserParams struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateUserParams struct {
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}
