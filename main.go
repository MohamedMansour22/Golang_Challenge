package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	Transaction "main.go/Transactions"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/Transactions", func(w http.ResponseWriter, r *http.Request) {
		transactions := Transaction.GetAllTransactions()
		json.NewEncoder(w).Encode(transactions)
	})

	r.Post("/Transactions/create", Transaction.CreateTransaction)

	http.ListenAndServe(":3000", r)
}
