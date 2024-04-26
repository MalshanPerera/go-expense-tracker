package routes

import (
	"github.com/MalshanPerera/go-expense-tracker/handlers"
	"github.com/labstack/echo"
)

type ExpenseRoute struct {
	V1      *echo.Group
	Handler *handlers.ExpenseHandler
}

func NewExpenseRoute(v1 *echo.Group, expenseHandler *handlers.ExpenseHandler) *ExpenseRoute {
	return &ExpenseRoute{
		V1:      v1,
		Handler: expenseHandler,
	}
}

func (route *ExpenseRoute) RegisterExpenseRoutes() {
	route.V1.POST("/expense", route.Handler.CreateExpense)
	route.V1.GET("/expense", route.Handler.GetExpense)
	route.V1.PUT("/expense", route.Handler.UpdateExpense)
	route.V1.DELETE("/expense", route.Handler.DeleteExpense)
}
