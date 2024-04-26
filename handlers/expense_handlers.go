package handlers

import (
	"net/http"

	"github.com/MalshanPerera/go-expense-tracker/services"
	"github.com/labstack/echo"
)

type ExpenseHandler struct {
	ExpenseService services.ExpenseServiceInterface
}

func NewExpenseHandler(expenseService services.ExpenseServiceInterface) *ExpenseHandler {
	return &ExpenseHandler{
		ExpenseService: expenseService,
	}
}

func (h *ExpenseHandler) GetExpense(c echo.Context) error {
	return c.String(http.StatusOK, "Get Handler")
}

func (h *ExpenseHandler) CreateExpense(c echo.Context) error {
	return c.String(http.StatusOK, "Create Handler")
}

func (h *ExpenseHandler) UpdateExpense(c echo.Context) error {
	return c.String(http.StatusOK, "Update Handler")
}

func (h *ExpenseHandler) DeleteExpense(c echo.Context) error {
	return c.String(http.StatusOK, "Delete Handler")
}
