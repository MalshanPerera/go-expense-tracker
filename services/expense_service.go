package services

import (
	"context"

	"github.com/MalshanPerera/go-expense-tracker/database/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ExpenseService struct {
	db      *pgxpool.Pool
	queries *sqlc.Queries
}

type ExpenseServiceInterface interface {
	GetExpenses(ctx context.Context, userPayload any) error
	CreateExpense(ctx context.Context, userPayload any) error
	UpdateExpense(ctx context.Context, userPayload any) error
	DeleteExpense(ctx context.Context, userPayload any) error
}

func NewExpenseService(db *pgxpool.Pool, queries *sqlc.Queries) ExpenseServiceInterface {
	return &ExpenseService{db: db, queries: queries}
}

func (c *ExpenseService) GetExpenses(ctx context.Context, userPayload any) error {
	return nil
}

func (c *ExpenseService) CreateExpense(ctx context.Context, userPayload any) error {
	return nil
}

func (c *ExpenseService) UpdateExpense(ctx context.Context, userPayload any) error {
	return nil
}

func (c *ExpenseService) DeleteExpense(ctx context.Context, userPayload any) error {
	return nil
}
