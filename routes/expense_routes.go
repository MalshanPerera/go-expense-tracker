package routes

import (
	"fmt"
	"net/http"

	"github.com/MalshanPerera/go-expense-tracker/handlers"
)

type ExpenseRoute struct {
	Handler *handlers.ExpenseHandler
}

func NewExpenseRoute(expenseHandler *handlers.ExpenseHandler) *ExpenseRoute {
	return &ExpenseRoute{
		Handler: expenseHandler,
	}
}

func (route *ExpenseRoute) RegisterExpenseRoutes() http.Handler {
	expenseHandlers := http.NewServeMux()

	handlersMap := map[string]func(w http.ResponseWriter, r *http.Request){
		"GET":    route.Handler.GetExpense,
		"POST":   route.Handler.CreateExpense,
		"PATCH":  route.Handler.UpdateExpense,
		"DELETE": route.Handler.DeleteExpense,
	}

	for method, handlerFunc := range handlersMap {
		handler := handlers.HandleFunc(handlerFunc)
		expenseHandlers.HandleFunc(fmt.Sprintf("%s /expense", method), handler)
	}

	return expenseHandlers
}
