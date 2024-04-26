package handlers

import (
	"net/http"

	"github.com/MalshanPerera/go-expense-tracker/services"
)

type ExpenseHandler struct {
	ExpenseService services.ExpenseServiceInterface
}

func NewExpenseHandler(expenseService services.ExpenseServiceInterface) *ExpenseHandler {
	return &ExpenseHandler{
		ExpenseService: expenseService,
	}
}

func (h *ExpenseHandler) GetExpense(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Get Handler"))
	if err != nil {
		return
	}
}

func (h *ExpenseHandler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Create Handler"))
	if err != nil {
		return
	}

}

func (h *ExpenseHandler) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Update Handler"))
	if err != nil {
		return
	}
}

func (h *ExpenseHandler) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Delete Handler"))
	if err != nil {
		return
	}
}
