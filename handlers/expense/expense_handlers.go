package expense_handlers

import (
	"net/http"
)

func Init() http.Handler {
	expenseHandlers := http.NewServeMux()

	expenseHandlers.HandleFunc("GET /expense", getExpense)
	expenseHandlers.HandleFunc("POST /expense", addExpense)
	expenseHandlers.HandleFunc("PATCH /expense", updateExpense)
	expenseHandlers.HandleFunc("DELETE /expense", deleteExpense)

	return expenseHandlers
}

func getExpense(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Get Handler"))
	if err != nil {
		return
	}
}

func addExpense(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Add Handler"))
	if err != nil {
		return
	}

}

func updateExpense(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Update Handler"))
	if err != nil {
		return
	}
}

func deleteExpense(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Delete Handler"))
	if err != nil {
		return
	}
}
